import React, { useState } from 'react';

const DeleteAccount: React.FC = () => {
    const [email, setEmail] = useState('');
    const [errorMessage, setErrorMessage] = useState('');

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setErrorMessage('');
        
        try {
            const response = await fetch('http://localhost:8000/delete', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ email }),
                credentials: 'include',
            });

            if (response.ok) {
                console.log('Account deletion successful');
                alert('アカウントが正常に削除されました。');
            } else {
                const errorData = await response.json();
                console.error('Account deletion failed:', errorData.message);
                setErrorMessage(errorData.message || 'アカウントの削除に失敗しました。');
            }
        } catch (error) {
            console.error('Network error:', error);
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
