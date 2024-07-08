import React from 'react';
import './App.css';
import { BrowserRouter as Router, Route, Routes, Navigate } from 'react-router-dom';
import Taskmanagement from './components/common/Taskmanagement';
import Signup from './components/auth/Signup';
import Signin from './components/auth/Signin';
import DeleteAccount from './components/auth/DeleteAccount';
import Home from './components/common/Home';
import Logout from './components/auth/Logout';
import AuthGuard from './components/common/AuthGuard';
import Profile from './components/user/Profile';

const App: React.FC = () => {
  return (
    <div className="App">
      <Router>
        <Routes>
          <Route path="/" element={<Taskmanagement />} />
          <Route path="/Signup" element={<Signup />} />
          <Route path="/Signin" element={<Signin />} />
          <Route path="/Delete" element={<DeleteAccount />} />
          <Route path="/Home" element={
            <AuthGuard>
              <Home />
            </AuthGuard>
          } />
          <Route path="/Logout" element={
            <AuthGuard>
              <Logout />
            </AuthGuard>}
          />
          <Route path="/:userId" element={
            <AuthGuard>
              <Profile />
            </AuthGuard>}
          />
          <Route path="*" element={<Navigate to="/" />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
