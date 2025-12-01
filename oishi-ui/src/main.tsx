import { QueryClient, QueryClientProvider, } from '@tanstack/react-query';
import React, { Suspense } from 'react';
import ReactDOM from 'react-dom/client';
import { RouterProvider, createRouter, NotFoundRoute } from '@tanstack/react-router'
import AuthProvider, { useAuth } from './Context/AuthContext';
import './index.css';
import { Route } from './Pages/__root.tsx'
import { routeTree } from './routeTree.gen';
import Loading from './components/Loading.tsx';


const queryClient = new QueryClient();

const notFoundRoute = new NotFoundRoute({
  getParentRoute: () => Route,
  component: () => "404 Not Found",
})
const router = createRouter({
  routeTree,
  notFoundRoute,
  defaultPreload: 'intent',
  context: {
    auth: undefined!
  }
})


declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router
  }
}
function InnerApp() {
  const auth = useAuth()
  return <RouterProvider router={router} context={{ auth }} />
}

const rootElement: HTMLElement | null = document.getElementById('root')
if (rootElement) {
  ReactDOM.createRoot(rootElement).render(
    <Suspense fallback={<Loading />}>
      <React.StrictMode>
        <AuthProvider>
          <QueryClientProvider client={queryClient}>
            <InnerApp />
          </QueryClientProvider>
        </AuthProvider>
      </React.StrictMode>
    </Suspense>
  );
} else {
  alert("ROOT ELEMENT not found")
}

