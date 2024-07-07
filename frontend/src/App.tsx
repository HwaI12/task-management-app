import React from 'react';
import './App.css';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import Register from './handlers/Register';
import Login from './handlers/Login';
import DeleteAccount from './handlers/DeleteAccount';
import Home from './components/Home';
import Logout from './handlers/Logout';

function App() {
  return (
    <div className="App">
      <Router>
        <nav>
          <Link to="/register">新規登録</Link> | <Link to="/login">ログイン</Link> | <Link to="/logout">ログアウト</Link>
        </nav>
        <Routes>
          <Route path="/register" element={<Register />} />
          <Route path="/login" element={<Login />} />
          <Route path="/delete" element={<DeleteAccount />} />
          <Route path="/home" element={<Home />} />
          <Route path="/logout" element={<Logout />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
