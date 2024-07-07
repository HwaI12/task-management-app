// components/Home.tsx

import React from 'react';
import { Navigate } from 'react-router-dom';

const isAuthenticated = (): boolean => {
    // Check if the user is authenticated
    return localStorage.getItem('authToken') !== null;
};

const Home: React.FC = () => {
    // Check authentication status
    if (!isAuthenticated()) {
        // If not authenticated, redirect to /login
        return <Navigate to="/login" replace />;
    }

    // If authenticated, render the Home component
    return (
        <div>
            <h2>Home Page</h2>
            <p>Welcome to the Home Page!</p>
        </div>
    );
};

export default Home;