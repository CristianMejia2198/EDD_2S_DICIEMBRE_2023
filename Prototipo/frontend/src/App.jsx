import { useState } from "react";
import { Route, Routes } from "react-router-dom";
import Login from "./pages/Login";
import Estudiante from "./pages/Estudiante";
import Tutor from "./pages/Tutor";
import Administrador from "./pages/Administrador";
import "./App.css";

function App() {
  const [count, setCount] = useState(0);

  return (
    <>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/principal/estudiante" element={<Estudiante />} />
        <Route path="/principal/tutor" element={<Tutor />} />
        <Route path="/principal/admin" element={<Administrador />} />
      </Routes>
    </>
  );
}

export default App;
