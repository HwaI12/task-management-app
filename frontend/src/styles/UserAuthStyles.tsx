// UserAuthStyles.tsx

import styled from 'styled-components';

export const Container = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background-color: #f1f1f1;
  padding: 20px;
  transition: background-color 0.3s ease-in-out;
  box-sizing: border-box;
`;

export const Form = styled.div`
  background: white;
  padding: 2rem;
  border-radius: 10px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 400px; /* 最大幅を設定 */
  border: 1px solid #d9dae2;
  box-sizing: border-box;
  transition: box-shadow 0.3s ease-in-out;
  margin: auto; /* 要素を水平方向に中央に配置 */
`;


export const Title = styled.h1`
  font-size: 1.5rem;
  margin-bottom: 1.5rem;
  color: #333;
  text-align: center;
`;

export const Label = styled.label`
  display: block;
  text-align: left;
  width: 100%;
  margin-bottom: 0.3rem;
  color: #333;
  font-size: 1rem;
`;

export const Input = styled.input`
  width: 100%;
  padding: 0.75rem;
  margin-bottom: 1rem;
  border: 1px solid #d9dae2;
  border-radius: 5px;
  font-size: 1rem;
  box-sizing: border-box;
  transition: border-color 0.3s ease-in-out;

  &:focus {
    border-color: #2e4d61;
  }
`;

export const Button = styled.button`
  width: 100%;
  padding: 0.75rem;
  margin-bottom: 1rem;
  border: none;
  border-radius: 5px;
  background-color: #eede77;
  color: #333;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s ease-in-out, transform 0.2s ease-in-out;

  &:hover {
    background-color: #e0d569;
  }

  &:disabled {
    background-color: #aaa;
    cursor: not-allowed;
  }
`;

export const LinkText = styled.p`
  color: #333;
  font-size: 0.875rem;
  text-align: right;
`;
