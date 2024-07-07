import React from 'react';
import './App.css';
import { BrowserRouter, Route, Routes, Link } from "react-router-dom";
import Register from './handlers/Register';
import Login from './handlers/Login';
import DeleteAccount from './handlers/DeleteAccount';

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Link to="/register">新規登録</Link> | <Link to="/login">ログイン</Link>
        <Routes>
          <Route path="/register" element={<Register />} />
          <Route path="/login" element={<Login />} />
          <Route path="/delete" element={<DeleteAccount />} />
        </Routes>
      </BrowserRouter>
    </div>
  )
}

export default App;