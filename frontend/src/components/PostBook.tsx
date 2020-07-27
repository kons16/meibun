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
        axios.post('http://localhost:8000/post_book',
            {'sentence': this.state.sentence, 'title': this.state.title, 'author': this.state.author, 'pages': this.state.pages},
            {withCredentials: true})
            .then((response) => {
                /// TODO
                console.log("ok");
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
                        <span className="label">名文</span>
                        <input type="text" name="sentence" onChange={this.onChange} />
                    </div>
                    <div>
                        <span className="label">本のタイトル</span>
                        <input type="text" name="title" onChange={this.onChange} />
                    </div>
                    <div>
                        <span className="label">ページ数</span>
                        <input type="text" name="pages" onChange={this.onChange} />
                    </div>
                    <div>
                        <span className="label">著者名</span>
                        <input type="text" name="author" onChange={this.onChange} />
                    </div>
                    <Button variant="contained" color="primary"　style={{ marginTop: 10, width: 100 }}
                            onClick={this.handleFormSubmit} >
                        追加する
                    </Button>
                </div>
                <Link to="/">ホームへ</Link>
            </div>
        );
    }
}

export default PostBook;
