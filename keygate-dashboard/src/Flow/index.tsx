import { useEffect, useState } from "react";
import ReactFlow, {
  Controls,
  Background,
  Node,
  useNodesState,
  Edge,
  useEdgesState,
  MarkerType,
  useReactFlow,
} from "reactflow";
import "reactflow/dist/style.css";
import Dagre from "@dagrejs/dagre";
import KeyNode from "./KeyNode";
import PermissionNode from "./PermissionNode";

const g = new Dagre.graphlib.Graph().setDefaultEdgeLabel(() => ({}));

interface Key {
  id: string;
  name: string;
  permissions: Permission[];
}

interface Permission {
  id: string;
  name: string;
}

const nodeTypes = {
  key: KeyNode,
  permission: PermissionNode,
};

const getLayoutedElements = (nodes: Node[], edges: Edge[]) => {
  if (nodes.length === 0) {
    return { nodes, edges };
  }

  g.setGraph({ rankdir: "TB", nodesep: 50, ranksep: 100 });

  edges.forEach((edge: Edge) => g.setEdge(edge.source, edge.target));
  nodes.forEach((node: Node) =>
    g.setNode(node.id, {
      label: node.id,
      width: 100,
      height: 100,
    })
  );

  Dagre.layout(g);

  return {
    nodes: nodes.map((node) => {
      const { x, y } = g.node(node.id);

      return { ...node, position: { x, y } };
    }),
    edges,
  };
};

function Flow() {
  const { fitView } = useReactFlow();
  const [keys, setKeys] = useState<Key[]>([]);
  const [permissions, setPermissions] = useState<Permission[]>([]);
  const [nodes, setNodes, onNodesChange] = useNodesState<Node[]>([]);
  const [edges, setEdges, onEdgesChange] = useEdgesState<Edge[]>([]);

  // Fetch keys from API
  useEffect(() => {
    fetch("http://localhost:8080/keys")
      .then((res) => {
        return res.json();
      })
      .then((res) => {
        const permissions = new Set<Permission>();

        res.data.forEach((key: Key) => {
          key.permissions.forEach((permission) => {
            permissions.add(permission);
          });
        });

        setKeys(res.data);
        setPermissions(Array.from(permissions));
      });
  }, []);

  // On keys change, update nodes
  useEffect(() => {
    const keyNodes: Node[] = keys.map((key) => {
      return {
        id: key.id,
        data: { label: key.name },
        position: { x: 0, y: 0 },
        type: "key",
      };
    });

    // Permissions is a Set of Permissions (they have a unique id)
    // For each key, we need to create an edge to each permission
    const edges: Edge[] = keys.flatMap((key) => {
      return Array.from(permissions).map((permission) => {
        return {
          id: `${key.id}-${permission.id}`,
          source: key.id,
          target: permission.id,
          animated: false,
          markerEnd: {
            type: MarkerType.ArrowClosed,
            width: 20,
            height: 20,
            color: "#b1b1b7",
          },
        };
      });
    });

    // Also, we need to create the actual nodes for each permission
    const permissionNodes: Node[] = Array.from(permissions).map(
      (permission) => {
        return {
          id: permission.id,
          data: { label: permission.name, color: "#b1b1b7" },
          position: { x: 0, y: 0 },
          type: "permission",
        };
      }
    );

    const layouted = getLayoutedElements(
      [...keyNodes, ...permissionNodes],
      edges
    );
    setNodes([...layouted.nodes]);
    setEdges([...layouted.edges]);
  }, [keys]);

  return (
    <div style={{ height: "100%" }}>
      <ReactFlow
        nodes={nodes}
        edges={edges}
        fitView
        nodeTypes={nodeTypes}
        onNodesChange={onNodesChange}
      >
        <Background />
        <Controls />
      </ReactFlow>
    </div>
  );
}

export default Flow;
