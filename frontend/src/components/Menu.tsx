import React, { Component } from 'react';
import axios from 'axios';
import Button from "@material-ui/core/Button";
import history from '../history';

interface State {
    isLoggedIn: boolean
    user: {
        id: number,
        name: string,
    }
}

class Menu extends Component<{}, State> {
    state: State = {
        isLoggedIn: false,
        user: {
            id: 0,
            name: "test",
        }
    };

    // GET / でログインしているならユーザー情報を取得する
    componentDidMount() {
        axios.get('http://localhost:8000/', {withCredentials: true})
            .then((response) => {
                const userData = response.data.User;
                if(userData != null){
                    this.setState({
                        isLoggedIn: true,
                        user: {
                            id: userData.ID,
                            name: userData.Name,
                        }
                    })
                }
            })
            .catch(() => {
                console.log("index fail");
            });
    }

    // POST /signout してcookieを削除
    handleSignout = () => {
        axios.post('http://localhost:8000/signout',
            {}, {withCredentials: true})
            .then((response) => {
                const name = response.data.Name;
                document.cookie = `${name}=; max-age=0`;
                window.location.reload();
            })
            .catch(() => {
                console.log("signout fail");
            });
    }

    // マイページ /users/:id に遷移
    handleToMyPage = () => {
        history.push(`/users/${this.state.user.id}`)
    }

    // ログイン /login に遷移
    handleToLogin = () => {
        history.push('/login')
    }

    // 新規登録 /signup に遷移
    handleToSignup = () => {
        history.push('/signup')
    }

    render() {
        return (
            <div className="Menu">
                Topページ<br/>
                {(() => {
                    if (this.state.isLoggedIn) {
                        return (
                            <div>
                                <span>{this.state.user.name}</span> <br/>
                                <Button variant="contained" color="primary"　style={{ marginTop: 10, width: 110 }}
                                        onClick={this.handleSignout} >
                                    ログアウト
                                </Button>

                                <Button variant="contained" color="primary"　style={{ marginTop: 10, marginLeft: 10, width: 120 }}
                                        onClick={this.handleToMyPage} >
                                    マイページへ
                                </Button>
                            </div>
                        )
                    } else {
                        return (
                            <div>
                                <Button onClick={this.handleToLogin}>ログイン</Button>
                                <Button onClick={this.handleToSignup}>新規登録</Button>
                            </div>
                        );
                    }
                })()}
            </div>
        );
    }
}

export default Menu;
