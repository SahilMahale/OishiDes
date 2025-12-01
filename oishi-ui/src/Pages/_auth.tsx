import { Outlet, createFileRoute, redirect } from '@tanstack/react-router'

export const Route = createFileRoute('/_auth')({
  beforeLoad: ({ context, location }) => {
    /* console.log("Auth from router context:", context.auth.Context) */
    //IDK WHY TS server is not catching the context passed but it is there
    if (!context.auth.Context.isLoggedIn) {
      throw redirect({
        to: '/Login',
        search: {
          redirect: location.href
        }
      })
    }
  },
  component: () => <Outlet />
})
