import React, { Component } from 'react';
import axios from "axios";
import Book from "./Book";
import history from "../history";
import Button from "@material-ui/core/Button";

interface MyHartPageState {
    books: any
    user: {
        id: number,
    }
}

// 自分はハートしたbook一覧を表示するページ
class MyHartPage extends Component<{}, MyHartPageState> {
    constructor(props: any) {
        super(props);
        this.state = {
            books: [],
            user: {
                id: 0,
            }
        }
    }

    componentDidMount() {
        axios.get('http://localhost:8000/get_my_harts', {withCredentials: true})
            .then((response) => {
                const myHartBooks: any[] = response.data.myHartBooks;
                if(myHartBooks != null){
                    this.setState({
                        books: myHartBooks,
                        user: {
                            id: response.data.ID
                        }
                    })
                }
            })
            .catch(() => {
                console.log("books get fail");
            });
    }

    // マイページ /users/:id に遷移。ログインしている自分のIDを渡す
    handleToMyPage = () => {
        history.push({
            pathname: `/users/${this.state.user.id}`
        })
    }

    render() {
        return (
            <div>
                <Button variant="contained" color="primary"　style={{ marginTop: 10, marginLeft: 10, width: 120 }}
                        onClick={this.handleToMyPage} >
                    マイページへ
                </Button>

                {(() => {
                    const bookItems: any = [];
                    this.state.books.forEach((key: any, index: number) => {
                        bookItems.push(
                            <Book
                                key={index}
                                id={this.state.books[index].ID}
                                sentence={this.state.books[index].Sentence}
                                title={this.state.books[index].Title}
                                author={this.state.books[index].Author}
                                pages={this.state.books[index].Pages}
                                harts={this.state.books[index].Harts}
                                myPageFlag={false}
                                hartFlag={true}
                                bookUserID={this.state.books[index].UserID}
                            />
                        )
                    });

                    if (this.state.books) {
                        return (
                            <div>
                                {bookItems}
                            </div>
                        )
                    }
                })()}
            </div>
        );
    }
}

export default MyHartPage;
