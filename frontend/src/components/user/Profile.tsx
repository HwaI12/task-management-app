import React from 'react';
import { useParams } from 'react-router-dom';
import Sidebar from '../common/Sidebar';
import { ContentContainer } from '../../styles/SidebarStyles';

const Profile: React.FC = () => {
    const { userId } = useParams<{ userId: string }>();

    return (
        <div>
            <ContentContainer>
            <Sidebar />
            <h2>Profile Page</h2>
            <p>Welcome, User {userId}!</p>
            </ContentContainer>

        </div>
    );
};

export default Profile;
