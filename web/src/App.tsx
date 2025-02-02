import { useState } from "react"
import { DisplayGraph } from "./DisplayGraph"

export default function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <div className="absolute -z-50">
        <DisplayGraph sometext={"Node " + count}></DisplayGraph>
      </div>
      <div>
        <button
          className="bg-slate-800 text-slate-200 m-3 p-2 rounded-md"
          onClick={() => setCount((count) => count + 1)}
        >
          count is {count}
        </button>
      </div>
    </>
  )
}
