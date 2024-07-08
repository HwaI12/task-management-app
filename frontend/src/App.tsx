import React from 'react';
import './App.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Taskmanagement from './components/Taskmanagement';
import Register from './handlers/Register';
import Login from './handlers/Login';
import DeleteAccount from './handlers/DeleteAccount';
import Home from './components/Home';
import Logout from './handlers/Logout';

function App() {
  return (
    <div className="App">
      <Router>
        <Routes>
          <Route path="/" element={<Taskmanagement />} />
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
