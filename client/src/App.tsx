import { Route, Routes } from "react-router-dom"
import Home from "./pages/Home"
import Join from "./pages/Join"
import CreateGame from "./pages/CreateGame"
import Game from "./pages/Game"
import NotFound from "./pages/NotFound"

function App() {
  return <Routes>
    <Route path="/" element={<Home />} />
    <Route path="/join" element={<Join />} />
    <Route path="/create" element={<CreateGame />} />
    <Route path="/game/:id" element={<Game />} />
    <Route path="*" element={<NotFound />} />
  </Routes>
}

export default App;
