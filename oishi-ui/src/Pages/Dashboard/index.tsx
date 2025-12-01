import { createFileRoute } from '@tanstack/react-router'
import { Card, CardHeader, CardDescription, CardContent, CardTitle } from '@/components/ui/card'
export const Route = createFileRoute('/Dashboard/')({
  component: Dashboard
})
function Dashboard() {
  return (
    <div className="bg-slate-900 min-h-screen py-2 px-2">
      <div className="container mx-auto rounded-md bg-slate-800 flex flex-col items-center ">
        <h2 className="font-sans text-slate-200 tracking-tight text-3xl text-center font-bold ">
          Dashboard Page
        </h2>
        <div className="container flex flex-row px-2 py-2">
          {
            <>
              <Card id={"1"}
                onClick={() => {
                  //    console.log("Table Created with: ", tableInfoRef.current, "With state: ", state)
                }}
                className="w-60 h-80 bg-gray-700 text-slate-300 cursor-pointer rounded-lg border-1 
      hover:ring-2 hover:ring-cyan-400 drop-shadow-2xl shadow-lg shadow-slate-900 hover:shadow-xl hover:shadow-cyan-700/60">
                <CardHeader className="relative items-center bg-gray-700 text-gray-100">
                  <CardTitle className="text-center">
                    {`Card 1`}
                  </CardTitle>
                  <CardDescription>
                    Table Description
                  </CardDescription>
                </CardHeader>
                <CardContent className="">
                  <div className="">
                    <h5>
                      CPU: {"4"}
                    </h5>
                    <h5>
                      Ram : {"5gb"}
                    </h5>
                  </div>
                </CardContent>
              </Card>
              <Card id={"1"}
                onClick={() => {
                  //    console.log("Table Created with: ", tableInfoRef.current, "With state: ", state)
                }}
                className="w-60 h-80 bg-gray-700 text-slate-300 cursor-pointer rounded-lg border-1 
      hover:ring-2 hover:ring-cyan-400 drop-shadow-2xl shadow-lg shadow-slate-900 hover:shadow-xl hover:shadow-cyan-700/60">
                <CardHeader className="relative items-center bg-gray-700 text-gray-100">
                  <CardTitle className="text-center">
                    {`Card 1`}
                  </CardTitle>
                  <CardDescription>
                    Table Description
                  </CardDescription>
                </CardHeader>
                <CardContent className="">
                  <div className="">
                    <h5>
                      CPU: {"4"}
                    </h5>
                    <h5>
                      Ram : {"5gb"}
                    </h5>
                  </div>
                </CardContent>
              </Card>
              <Card id={"1"}
                onClick={() => {
                  //    console.log("Table Created with: ", tableInfoRef.current, "With state: ", state)
                }}
                className="w-60 h-80 bg-gray-700 text-slate-300 cursor-pointer rounded-lg border-1 
      hover:ring-2 hover:ring-cyan-400 drop-shadow-2xl shadow-lg shadow-slate-900 hover:shadow-xl hover:shadow-cyan-700/60">
                <CardHeader className="relative items-center bg-gray-700 text-gray-100">
                  <CardTitle className="text-center">
                    {`Card 1`}
                  </CardTitle>
                  <CardDescription>
                    Table Description
                  </CardDescription>
                </CardHeader>
                <CardContent className="">
                  <div className="">
                    <h5>
                      CPU: {"4"}
                    </h5>
                    <h5>
                      Ram : {"5gb"}
                    </h5>
                  </div>
                </CardContent>
              </Card>
            </>
          }
        </div>
      </div>
    </div>
  );
}
