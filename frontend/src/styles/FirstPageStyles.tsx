// LoginPageStyles.tsx

import styled from 'styled-components';

export const PageContainer = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  height: 100vh;
  background-color: #fff;
  padding: 20px;
`;

export const LogoContainer = styled.div`
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 5rem;
  font-weight: bold;
  color: #000;
`;

export const FormContainer = styled.div`
  flex: 1;
  max-width: 400px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
`;

export const Title = styled.h1`
  font-size: 2rem;
  margin-bottom: 1.5rem;
  color: #333;
  text-align: center;
`;

export const SubTitle = styled.p`
  font-size: 1.25rem;
  margin-bottom: 1.5rem;
  color: #333;
  text-align: center;
`;

export const Button = styled.button`
  width: 100%;
  padding: 0.75rem;
  margin-bottom: 1rem;
  border: none;
  border-radius: 5px;
  background-color: #e0e0e0;
  color: #333;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s ease;

  &:hover {
    background-color: #ccc;
  }
`;

export const LinkText = styled.p`
  color: #333;
  font-size: 0.875rem;
  text-align: center;

  a {
    color: #007bff;
    text-decoration: none;

    &:hover {
      text-decoration: underline;
    }
  }
`;
