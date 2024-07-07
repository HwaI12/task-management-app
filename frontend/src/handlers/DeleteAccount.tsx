import React, { useState } from 'react';

const DeleteAccount: React.FC = () => {
    const [email, setEmail] = useState('');
    const [errorMessage, setErrorMessage] = useState('');

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        try {
            const response = await fetch('http://localhost:8000/delete', {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ email }),
            });

            if (response.ok) {
                console.log('退会成功');
                alert('アカウントが正常に削除されました。');
                // 退会成功時の追加の処理をここに追加
            } else {
                console.error('退会失敗:', response.status);
                // 退会失敗時の処理をここに追加
                setErrorMessage('アカウントの削除に失敗しました。');
            }
        } catch (error) {
            console.error('ネットワークエラー:', error);
            // ネットワークエラー時の処理をここに追加
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
