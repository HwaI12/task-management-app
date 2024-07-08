import React, { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { Link } from 'react-router-dom';
import { PageContainer, LogoContainer, FormContainer, Title, SubTitle, Button, LinkText } from '../styles/FirstPageStyles';

const isAuthenticated = (): boolean => {
    return localStorage.getItem('authToken') !== null;
};

const TaskManagement: React.FC = () => {
    const navigate = useNavigate();

    useEffect(() => {
        if (!isAuthenticated()) {
            navigate('/login');
        }
    }, [navigate]);

    return (
        <PageContainer>
            <LogoContainer>X</LogoContainer>
            <FormContainer>
                <Title>Happening now</Title>
                <SubTitle>Join today.</SubTitle>
                <Button>Sign up with Google</Button>
                <Button>Sign up with Apple</Button>
                <Button>Create account</Button>
                <LinkText>
                    <p>Already have an account?</p>
                    <Link to="/login">Sign in</Link>
                </LinkText>
            </FormContainer>
        </PageContainer>
    );
};

export default TaskManagement;
