import ApolloClient from 'apollo-client';
import { InMemoryCache } from 'apollo-cache-inmemory';
import builder from './../commons/builder';
import resolvers from './resolvers';
import defaults from './defaults';
import { getURL } from './../commons/api-url';
import { WebSocketLink } from './ws';
import { createHttpLink } from 'apollo-link-http';
import { split, ApolloLink } from 'apollo-link';
import { getMainDefinition } from 'apollo-utilities';
import { withClientState } from 'apollo-link-state';
import { setContext } from 'apollo-link-context';
import { createTransformerLink } from 'apollo-client-transform';
import { getLinkTransformers } from './transformers';

export function createApolloClient() {
  const graphqlApiUrl = getURL(
    process.env.REACT_APP_LOCAL_API ? 'graphqlApiUrlLocal' : 'graphqlApiUrl',
  );

  const httpLink = createHttpLink({ uri: graphqlApiUrl });
  const authLink = setContext((_, { headers }) => {
    return {
      headers: {
        ...headers,
        authorization: builder.getBearerToken() || null,
      },
    };
  });
  const authHttpLink = authLink.concat(httpLink);

  const wsLink = new WebSocketLink({
    uri: getURL('subscriptionsApiUrl'),
    options: {
      reconnect: true,
    },
  });

  const cache = new InMemoryCache();

  const stateLink = withClientState({
    cache,
    defaults,
    resolvers,
  });

  const linkTransformers = getLinkTransformers();
  const transformerLink = createTransformerLink(linkTransformers);
  const enhancedAuthHttpLink = transformerLink.concat(authHttpLink);

  const link = split(
    ({ query }) => {
      const { kind, operation } = getMainDefinition(query);
      return kind === 'OperationDefinition' && operation === 'subscription';
    },
    wsLink,
    enhancedAuthHttpLink,
  );

  const client = new ApolloClient({
    uri: graphqlApiUrl,
    cache,
    link: ApolloLink.from([stateLink, link]),
    connectToDevTools: true,
  });

  return client;
}
