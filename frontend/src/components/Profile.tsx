// components/Profile.tsx

// 自分のプロフィールページ
// ルートはその人のユーザーIDにしたい
import React from 'react';
import { Navigate } from 'react-router-dom';

const isAuthenticated = (): boolean => {
    // Check if the user is authenticated
    return localStorage.getItem('authToken') !== null;
};

const Profile: React.FC = () => {
    // Check authentication status
    if (!isAuthenticated()) {
        // If not authenticated, redirect to /login
        return <Navigate to="/login" replace />;
    }

    // If authenticated, render the Profile component
    return (
        <div>
            <h2>Profile Page</h2>
            <p>Welcome to the Profile Page!</p>
        </div>
    );
};

export default Profile;