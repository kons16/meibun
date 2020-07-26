import React, { Component } from 'react';
import { Link } from "react-router-dom";
import axios from 'axios';
import Button from '@material-ui/core/Button';

interface State {
    sentence?: string
    title?: string
    author?: string
    pages?: number
}

// 名文を新規登録する画面
class PostBook extends Component<{}, State> {
    state: State = {
        sentence: "",
        title: "",
        author: "",
        pages: 0,
    };

    componentDidMount() {
    }

    onChange = (e: any) => {
        this.setState({
            [e.target.name]: e.target.value,
        });
    }

    // 名文情報をpostする
    handleFormSubmit = () => {
        axios.post('http://localhost:8000/new_book',
            {'sentence': this.state.sentence, 'title': this.state.title, 'author': this.state.author, 'pages': this.state.pages})
            .then((response) => {
            })
            .catch(() => {
                console.log("post fail");
            });
    }

    render() {
        return (
            <div>
                <div id="form">
                    <div>
                        <span className="label">名前</span>
                        <input type="text" name="name" onChange={this.onChange} />
                    </div>
                    <div>
                        <span className="label">メールアドレス</span>
                        <input type="text" name="email" onChange={this.onChange} />
                    </div>
                    <div>
                        <span className="label">パスワード</span>
                        <input type="password" name="password" onChange={this.onChange} />
                    </div>
                    <Button variant="contained" color="primary"　style={{ marginTop: 10, width: 100 }}
                            onClick={this.handleFormSubmit} >
                        新規登録
                    </Button>
                </div>
                <Link to="/">ホームへ</Link>
            </div>
        );
    }
}

export default PostBook;
