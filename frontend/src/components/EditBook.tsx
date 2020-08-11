import React, { Component } from 'react';
import Button from "@material-ui/core/Button";
import history from "../history";

interface EditBookProps {
    location: any
}

interface EditBookState {
    book: {
        id: number
        sentence: string
        title: string
        author: string
        pages: number
        harts: number
        myPageFlag: boolean
        hartFlag: boolean
        bookUserID: number
    }
}

// Book編集画面のコンポーネント
class EditBook extends Component<EditBookProps, EditBookState> {
    constructor(props: any) {
        super(props);
        this.state = {
            book: {
                id: this.props.location.state.Book.id,
                sentence: this.props.location.state.Book.sentence,
                title: this.props.location.state.Book.title,
                author: this.props.location.state.Book.author,
                pages: this.props.location.state.Book.pages,
                harts: this.props.location.state.Book.harts,
                myPageFlag: this.props.location.state.Book.myPageFlag,
                hartFlag: this.props.location.state.Book.hartFlag,
                bookUserID: this.props.location.state.Book.bookUserID
            }
        };
    }

    // マイページ /users/:id に遷移。ログインしている自分のIDを渡す
    handleToMyPage = () => {
        history.push({
            pathname: `/users/${this.state.book.bookUserID}`
        })
    }

    render() {
        return (
            <div>
                <div id="form">
                    <div>
                        <span className="label">名文</span>
                        <input type="text" name="sentence" value={this.state.book.sentence} />
                    </div>
                    <div>
                        <span className="label">本のタイトル</span>
                        <input type="text" name="title" value={this.state.book.title} />
                    </div>
                    <div>
                        <span className="label">ページ数</span>
                        <input type="text" name="pages" value={this.state.book.pages} />
                    </div>
                    <div>
                        <span className="label">著者名</span>
                        <input type="text" name="author" value={this.state.book.author} />
                    </div>
                    <Button variant="contained" color="primary"　style={{ marginTop: 10, width: 100 }}>
                        追加する
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
