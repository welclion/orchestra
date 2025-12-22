// frontend/src/components/ProjectList.js
import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

export default function ProjectList() {
  const [projects, setProjects] = useState([]);
  const [newProjectName, setNewProjectName] = useState('');
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  // Хуки вызываются ВСЕГДА — до любых условий
  const token = localStorage.getItem('token');

  useEffect(() => {
    // Проверка токена — внутри useEffect, а не снаружи
    if (!token) {
      navigate('/login');
      return;
    }

    const API_BASE = 'http://localhost:8080';

    const fetchProjects = async () => {
      try {
        const res = await fetch(`${API_BASE}/projects`, {
          headers: { Authorization: `Bearer ${token}` }
        });
        if (!res.ok) throw new Error('Не удалось загрузить проекты');
        const data = await res.json();
        setProjects(data.projects || []);
      } catch (err) {
        console.error(err);
        alert(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchProjects();
  }, [navigate, token]); // зависимости

  const createProject = async () => {
    if (!token || !newProjectName.trim()) return;
    try {
      const API_BASE = 'http://localhost:8080';
      const res = await fetch(`${API_BASE}/projects`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`
        },
        body: JSON.stringify({ name: newProjectName })
      });
      if (!res.ok) throw new Error('Не удалось создать проект');
      setNewProjectName('');
      // Перезагрузить список
      const res2 = await fetch(`${API_BASE}/projects`, {
        headers: { Authorization: `Bearer ${token}` }
      });
      const data = await res2.json();
      setProjects(data.projects || []);
    } catch (err) {
      alert(err.message);
    }
  };

  // Рендер — с проверкой токена
  if (!token) {
    return null; // или редирект уже сделан в useEffect
  }

  return (
    <div style={{ padding: '20px', maxWidth: '600px', margin: '0 auto' }}>
      <h2>Мои проекты</h2>
      
      <div style={{ marginBottom: '20px' }}>
        <input
          type="text"
          placeholder="Название проекта"
          value={newProjectName}
          onChange={(e) => setNewProjectName(e.target.value)}
          style={{ padding: '8px', width: '70%' }}
        />
        <button
          onClick={createProject}
          style={{ padding: '8px 16px', marginLeft: '10px' }}
        >
          Создать
        </button>
      </div>

      {loading ? (
        <p>Загрузка...</p>
      ) : projects.length === 0 ? (
        <p>У вас пока нет проектов.</p>
      ) : (
        <ul>
          {projects.map((p) => (
            <li key={p.id} style={{ padding: '8px', borderBottom: '1px solid #eee' }}>
              <strong>{p.name}</strong> — {p.status}
            </li>
          ))}
        </ul>
      )}

      <button
        onClick={() => {
          localStorage.removeItem('token');
          navigate('/login');
        }}
        style={{ marginTop: '20px' }}
      >
        Выйти
      </button>
    </div>
  );
}