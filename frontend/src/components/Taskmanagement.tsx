// TaskManagement.tsx

import React from 'react';
import { useNavigate } from 'react-router-dom';
import { PageContainer, FormContainer, Title, SubTitle, Button1, Button2} from '../styles/FirstPageStyles';

const TaskManagement: React.FC = () => {
    const navigate = useNavigate();

    const handleCreateAccount = () => {
        navigate('/register');
    };

    const handleSignIn = () => {
        navigate('/login');
    };

    return (
        <PageContainer>
            <FormContainer>
                <Title>時間を<br></br>有効活用しよう</Title>
                <SubTitle>AIを活用したワークスペースで成果物を共有しましょう</SubTitle>
                <Button1 onClick={handleCreateAccount}>Create account</Button1>
                <p>すでにアカウントをお持ちですか？</p>
                <Button2 onClick={handleSignIn}>Sign in</Button2>
            </FormContainer>
        </PageContainer>
    );
};

export default TaskManagement;
