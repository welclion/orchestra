// frontend/src/App.js
import React from 'react';
import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';
import Register from './components/Register';
import Login from './components/Login';
import ProjectList from './components/ProjectList';

function App() {
  return (
    <BrowserRouter>
      <div style={{ padding: '20px' }}>
        <nav>
          <Link to="/" style={{ marginRight: '16px' }}>Регистрация</Link>
          <Link to="/login" style={{ marginRight: '16px' }}>Вход</Link>
          <Link to="/projects">Проекты</Link>
        </nav>
        <Routes>
          <Route path="/" element={<Register />} />
          <Route path="/login" element={<Login />} />
          <Route path="/projects" element={<ProjectList />} />
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;