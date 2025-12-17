// frontend/src/App.js
import React, { useState } from 'react';

function App() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');

  const handleRegister = async (e) => {
    e.preventDefault();
    setMessage('Отправка...');

    try {
      const response = await fetch('http://localhost:8080/auth/register', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      });

      const data = await response.json();

      if (response.ok) {
        setMessage(`✅ Успешно! ID: ${data.id}`);
      } else {
        // Ошибка от сервера (например, "email уже существует")
        setMessage(`❌ Ошибка: ${data.error || JSON.stringify(data)}`);
      }
    } catch (err) {
      setMessage(`💥 Ошибка сети: ${err.message}`);
    }
  };

  return (
    <div style={{ padding: '20px', maxWidth: '500px', margin: '0 auto', fontFamily: 'sans-serif' }}>
      <h1>ORCHESTRA — Регистрация</h1>
      <form onSubmit={handleRegister}>
        <div style={{ marginBottom: '12px' }}>
          <label htmlFor="email">Email:</label>
          <input
            id="email"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
            style={{ width: '100%', padding: '8px', marginTop: '4px' }}
          />
        </div>

        <div style={{ marginBottom: '12px' }}>
          <label htmlFor="password">Пароль (6+ символов):</label>
          <input
            id="password"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            style={{ width: '100%', padding: '8px', marginTop: '4px' }}
          />
        </div>

        <button
          type="submit"
          style={{
            padding: '10px 20px',
            backgroundColor: '#4CAF50',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer',
          }}
        >
          Зарегистрироваться
        </button>
      </form>

      {message && (
        <div
          style={{
            marginTop: '16px',
            padding: '10px',
            backgroundColor: message.includes('Успешно') ? '#e8f5e9' : '#ffebee',
            borderLeft: `4px solid ${message.includes('Успешно') ? '#4CAF50' : '#f44336'}`,
          }}
        >
          {message}
        </div>
      )}
    </div>
  );
}

export default App;