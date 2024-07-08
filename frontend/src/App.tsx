import React from 'react';
import './App.css';
import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
import Taskmanagement from './components/Taskmanagement';
import Signup from './handlers/Signup';
import Signin from './handlers/Signin';
import DeleteAccount from './handlers/DeleteAccount';
import Home from './components/Home';
import Logout from './handlers/Logout';
import AuthGuard from './components/AuthGuard';

const App: React.FC = () => {
  return (
    <div className="App">
      <Router>
        <Routes>
          <Route path="/" element={<Taskmanagement />} />
          <Route path="/Signup" element={<Signup />} />
          <Route path="/Signin" element={<Signin />} />
          <Route path="/delete" element={<DeleteAccount />} />
          <Route path="/home" element={
            <AuthGuard>
              <Home />
            </AuthGuard>
          } />
          <Route path="/logout" element={
            <AuthGuard>
              <Logout />
            </AuthGuard>}
          />
          <Route path="*" element={<Navigate to="/" />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
