import React, { Component } from 'react';
import axios from "axios";
import Button from "@material-ui/core/Button";

interface BookProps {
    id: number
    sentence: string
    title: string
    author: string
    pages: number
    harts: number
    myPageFlag: boolean
}

interface BookState {
    id: number
    sentence: string
    title: string
    author: string
    pages: number
    harts: number
    myPageFlag: boolean
}

// Book自体のコンポーネント
class Books extends Component<BookProps, BookState> {
    constructor(props: any) {
        super(props);
        this.state = {
            id: props.id,
            sentence: props.sentence,
            title: props.title,
            author: props.author,
            pages: props.pages,
            harts: props.harts,
            myPageFlag: props.myPageFlag
        };
    }

    // bookを削除する
    handleDeleteBook = () => {
        axios.post('http://localhost:8000/delete_book',
            {'bookID': this.state.id},
            {withCredentials: true})
            .then((response) => {
                window.location.reload();
            })
            .catch(() => {
                console.log("delete fail");
            });
    }

    componentDidMount() {
    }

    render() {
        return (
            <div id="book-component">
                <div id="book-sentence">{this.state.sentence}</div>
                <div id="book-title">『{this.state.title}』</div>
                <div id="book-author">{this.state.author}</div>
                <div id="book-pages">p.{this.state.pages}</div>
                <div id="book-harts">♡{this.state.harts}</div>
                {this.state.myPageFlag &&
                    <div>
                        <Button variant="contained" color="primary"　style={{ marginTop: 10, width: 10 }}
                                onClick={this.handleDeleteBook} >
                            ☓
                        </Button>
                    </div>
                }
            </div>
        );
    }
}

export default Books;
