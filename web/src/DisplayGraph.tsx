import { useEffect } from "react"
import Graph from "graphology"
import { SigmaContainer, useLoadGraph } from "@react-sigma/core"
import "@react-sigma/core/lib/react-sigma.min.css"

const sigmaStyle = { height: "100vh", width: "100vw" }

// Component that load the graph
export const LoadGraph = ({ nodename }: { nodename: string }) => {
  const loadGraph = useLoadGraph()

  useEffect(() => {
    const graph = new Graph()
    graph.addNode("first", {
      x: 0,
      y: 0,
      size: 15,
      label: nodename,
      color: "#FA4F40",
    })
    graph.addNode("second", {
      x: 0.1,
      y: 3,
      size: 15,
      label: nodename,
      color: "#FA4F40",
    })
    loadGraph(graph)
  }, [loadGraph, nodename])

  return null
}

// Component that display the graph
export const DisplayGraph = ({ sometext }: { sometext: string }) => {
  return (
    <>
      <SigmaContainer style={sigmaStyle}>
        <LoadGraph nodename={sometext} />
      </SigmaContainer>
    </>
  )
}
