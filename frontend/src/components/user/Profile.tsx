import React from 'react';
import { useParams } from 'react-router-dom';

const Profile: React.FC = () => {
    const { userId } = useParams<{ userId: string }>();

    return (
        <div>
            <h2>Profile Page</h2>
            <p>Welcome, User {userId}!</p>
        </div>
    );
};

export default Profile;
