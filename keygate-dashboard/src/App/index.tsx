import { ReactFlowProvider } from "reactflow";
import Flow from "../Flow";

import "./App.css";

function App() {
  return (
    <div className="App">
      <header className="App-header">Keygate</header>
      <ReactFlowProvider>
        <Flow />
      </ReactFlowProvider>
    </div>
  );
}

export default App;
