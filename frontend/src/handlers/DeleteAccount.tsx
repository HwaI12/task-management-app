import React, { useState } from 'react';

// emailを使ってアカウント削除するのではなく、認証済みのユーザーのみがアカウントを削除できるようにしたい
// そのため、認証済みのユーザーのみがアカウント削除リクエストを送信できるように、認証済みのユーザーのみがこのページにアクセスできるようにする
// そのためには、認証済みのユーザーのみがこのページにアクセスできるようにするための認証機能を実装する必要がある

// アカウントを削除したら、そのユーザーのデータを削除するだけでなく、そのユーザーが作成したデータも削除する必要がある
// そのため、アカウント削除リクエストを受け取ったら、そのユーザーが作成したデータを削除する処理を実装する必要がある
// そのためには、データベースに保存されているデータを削除する処理を実装する必要がある

// ボタンを押したら/loginにリダイレクトするようにしたい
// そのためには、リダイレクト機能を実装する必要がある

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
