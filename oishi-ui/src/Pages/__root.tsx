import { AppContext } from "@/Context/AuthContext"
import Navbar from "@/components/Navbar"
import { createRootRouteWithContext } from "@tanstack/react-router"
import { TanStackRouterDevtools } from "@tanstack/router-devtools"
const isProd = import.meta.env.VITE_DEPLOY_ENV === "PROD"

interface MyRouterContext {
  auth: AppContext
}
export const Route = createRootRouteWithContext<MyRouterContext>()({
  component: () => (
    <>
      <Navbar />
      {!isProd && (
        <TanStackRouterDevtools position="bottom-right" initialIsOpen={false} />
      )}
    </>
  ),
  notFoundComponent: () => { return (<p className="text-red-500">Root component not Sound</p>) }
})


