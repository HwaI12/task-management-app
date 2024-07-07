import React, { useState } from 'react';

const DeleteAccount: React.FC = () => {
    const [email, setEmail] = useState('');
    const [errorMessage, setErrorMessage] = useState('');

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        try {
            const response = await fetch('http://localhost:8000/delete', {
                method: 'POST', // Changed to POST to match backend
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ email }),
                credentials: 'include', // Add this line to include cookies
            });
    
            if (response.ok) {
                console.log('アカウント削除成功');
                alert('アカウントが正常に削除されました。');
            } else {
                const errorData = await response.json();
                console.error('アカウント削除失敗:', errorData.message);
                setErrorMessage(errorData.message || 'アカウントの削除に失敗しました。');
            }
        } catch (error) {
            console.error('ネットワークエラー:', error);
            setErrorMessage('ネットワークエラーが発生しました。');
        }
    };

    return (
        <div>
            <h2>退会</h2>
            {errorMessage && <div style={{ color: 'red' }}>{errorMessage}</div>}
            <form onSubmit={handleSubmit}>
                <input
                    type="email"
                    placeholder="Email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                />
                <button type="submit">退会</button>
            </form>
        </div>
    );
};

export default DeleteAccount;
