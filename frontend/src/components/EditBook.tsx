import React, { Component } from 'react';
import Button from "@material-ui/core/Button";
import history from "../history";
import axios from "axios";

interface EditBookProps {
    location: any
}

interface EditBookState {
    id?: number
    sentence?: string
    title?: string
    author?: string
    pages?: number
    bookUserID?: number
}

// Book編集画面のコンポーネント
class EditBook extends Component<EditBookProps, EditBookState> {
    constructor(props: any) {
        super(props);
        this.state = {
            id: this.props.location.state.Book.id,
            sentence: this.props.location.state.Book.sentence,
            title: this.props.location.state.Book.title,
            author: this.props.location.state.Book.author,
            pages: this.props.location.state.Book.pages,
            bookUserID: this.props.location.state.Book.bookUserID
        };
    }

    onChange = (e: any) => {
        // console.log(e.target.name, e.target.value); => pages 41
        this.setState({
            [e.target.name]: e.target.value
        });
    }

    // 更新したbook情報をpostする
    submitEditBook =() => {
        axios.post('http://localhost:8000/update_book',
            {'ID': this.state.id, 'sentence': this.state.sentence, 'title': this.state.title, 'author': this.state.author,
                'pages': this.state.pages, 'userID': this.state.bookUserID},
            {withCredentials: true})
            .then((response) => {
                // マイページに遷移
                history.push({
                    pathname: `/users/${this.state.bookUserID}`
                });
            })
            .catch(() => {
                console.log("post fail");
            });
    }

    // マイページ /users/:id に遷移。ログインしている自分のIDを渡す
    handleToMyPage = () => {
        history.push({
            pathname: `/users/${this.state.bookUserID}`
        })
    }

    render() {
        return (
            <div>
                <div id="form">
                    <div>
                        <span className="label">名文</span>
                        <input type="text" name="sentence" value={this.state.sentence} onChange={this.onChange} />
                    </div>
                    <div>
                        <span className="label">本のタイトル</span>
                        <input type="text" name="title" value={this.state.title} onChange={this.onChange} />
                    </div>
                    <div>
                        <span className="label">ページ数</span>
                        <input type="text" name="pages" value={this.state.pages} onChange={this.onChange} />
                    </div>
                    <div>
                        <span className="label">著者名</span>
                        <input type="text" name="author" value={this.state.author} onChange={this.onChange} />
                    </div>
                    <Button variant="contained" color="primary"　style={{ marginTop: 10, width: 100 }} onClick={this.submitEditBook}>
                        編集
                    </Button>
                </div>

                <Button variant="contained" color="primary"　style={{ marginTop: 10, marginLeft: 10, width: 120 }}
                        onClick={this.handleToMyPage} >
                    マイページへ
                </Button>
            </div>
        );
    }
}

export default EditBook;
