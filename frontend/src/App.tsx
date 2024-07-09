import React from 'react';
import './App.css';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Taskmanagement from './components/common/Taskmanagement';
import Signup from './components/auth/Signup';
import Signin from './components/auth/Signin';
import DeleteAccount from './components/auth/DeleteAccount';
import Home from './components/common/Home';
import Logout from './components/auth/Logout';
import AuthGuard from './components/common/AuthGuard';
import Profile from './components/user/Profile';
import CreateTask from './components/user/CreateTask';
import NotFound from './components/common/NotFound';
import Notifications from './components/user/Notifications';
import ViewTask from './components/user/ViewTask';

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
          <Route path="/:userId/addtask" element={
            <AuthGuard requireMatch={true}>
              <CreateTask />
            </AuthGuard>}
          />
          <Route path="/Notifications" element={
            <AuthGuard>
              <Notifications />
            </AuthGuard>
          } />
          <Route path="/:userId/task/:taskId" element={
            <AuthGuard>
              <ViewTask />
            </AuthGuard>
          } />
          <Route path="*" element={<NotFound />} /> {/* 404ページのルートを追加 */}
        </Routes>
      </Router>
    </div>
  );
}

export default App;
