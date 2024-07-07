import React from 'react';
import './App.css';
import { BrowserRouter, Route, Routes, Link } from "react-router-dom";
import Register from './handlers/Register';
import Login from './handlers/Login';
import DeleteAccount from './handlers/DeleteAccount';
import Home from './components/Home';

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Link to="/register">新規登録</Link> | <Link to="/login">ログイン</Link>
        <Routes>
          <Route path="/register" element={<Register />} />
          <Route path="/login" element={<Login />} />
          <Route path="/delete" element={<DeleteAccount />} />
          <Route path="/home" element={<Home />} />
        </Routes>
      </BrowserRouter>
    </div>
  )
}

export default App;