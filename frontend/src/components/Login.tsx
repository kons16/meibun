import React, { Component } from 'react';
import axios from 'axios';

interface State {
    email?: string
    password?: string
}

// GET /signin でログイン済みなら / にリダイレクトさせ、未ログインならログインフォームを表示させる
class Login extends Component<{}, State> {
    state: State = {
        email: "",
        password: ""
    };

    // ログイン済みかどうかチェック
    componentDidMount() {
        console.log(localStorage.getItem('meibun_token'));
    }

    onChange = (e: any) => {
        this.setState({
            [e.target.name]: e.target.value,
        });
    }

    // formのsubmitが呼ばれたときの関数  /signin に POSTする
    handleFormSubmit = () => {
        axios.post('http://localhost:8000/signin',
            {'email': this.state.email, 'password': this.state.password})
            .then((response) => {
                const token = response.data;
                localStorage.setItem('meibun_token', token);
            })
            .catch(() => {
                console.log("login fail");
            });
    }

    render() {
        return (
            <form onSubmit={this.handleFormSubmit}>
                <div>
                    <span className="label">メールアドレス</span>
                    <input type="text" name="email" onChange={this.onChange} />
                </div>
                <div>
                    <span className="label">パスワード</span>
                    <input type="password" name="password" onChange={this.onChange} />
                </div>
                <button type="submit">ログイン</button>
            </form>
        );
    }
}

export default Login;
