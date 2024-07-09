import React from 'react';
import Sidebar from '../common/Sidebar';
import { ContentContainer } from '../../styles/SidebarStyles';

const Notifications: React.FC = () => {
    return (
        <div>
            <Sidebar />
            <ContentContainer>
                <h1>Notifications Page</h1>
                <p>追加予定</p>
            </ContentContainer>
        </div>
    );
};

export default Notifications;
