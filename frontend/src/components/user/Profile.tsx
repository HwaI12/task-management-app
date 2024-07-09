import React from 'react';
import { useParams } from 'react-router-dom';
// import Header from '../common/Header';
// import { ContentContainer } from '../../styles/HeaderStyles';

const Profile: React.FC = () => {
    const { userId } = useParams<{ userId: string }>();

    return (
        <div>
            {/* <Header />
            <ContentContainer> */}
            <h2>Profile Page</h2>
            <p>Welcome, User {userId}!</p>
            {/* </ContentContainer> */}

        </div>
    );
};

export default Profile;
