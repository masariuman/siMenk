import React, { Component } from "react";
import Footer from "../warudo/Footer";
import DarkMode from "../warudo/DarkMode";

import Loading from "../warudo/Loading";

class LandingPage extends Component {
    constructor(props) {
        super(props);
        this.state = {
            loading: true
        };
    }

    getData() {
        this.setState({
            loading: false
        });
    }

    componentDidMount() {
        this.getData();
    }

    render() {
        return (
            this.state.loading === true ? <Loading /> :
            <div className="content-w">
                {/* title content */}
                <div className="top-bar color-scheme-transparent masariuman-height103px">
                    <div className="top-menu-controls masariuman-marginleft30px">
                        <div className="icon-w top-icon masariuman-titlecontent" id="test">
                        <div className="os-icon os-icon-layout"></div>
                        </div>
                        <div className="masariuman-textleft">
                            <span className="masariuman-bold">Dashboard</span> <br/>
                            <small>Manajemen Dashboard</small>
                        </div>
                    </div>
                    <div className="top-menu-controls">
                        <button className="mr-2 mb-2 btn btn-outline-primary" type="button" id="petunjuk"><i className="batch-icon-bulb-2"></i> PETUNJUK <i className="batch-icon-bulb"></i></button>
                    </div>
                </div>
                <ul className="breadcrumb">
                    <li className="breadcrumb-item">
                        <a>Dashboard</a>
                    </li>
                    <li className="breadcrumb-item">
                        <span>Dashboard</span>
                    </li>
                </ul>

                {/* content */}
                <div className="content-i masariuman-minheight100vh">
                        <div className="content-box">
                            <div className="row">
                                <div className="col-sm-12">
                                    <div className="element-wrapper row">
                                        <div className="col-sm-3">
                                            <a className="element-box el-tablo centered trend-in-corner padded bold-label">
                                                <div className="value">
                                                    0
                                                </div>
                                                <div className="label">
                                                    Total Posts
                                                </div>
                                            </a>
                                        </div>
                                        <div className="col-sm-2">
                                            <a className="element-box el-tablo centered trend-in-corner padded bold-label masariuman_colorGreen">
                                                <div className="value">
                                                    0
                                                </div>
                                                <div className="label">
                                                    Total Novels
                                                </div>
                                            </a>
                                        </div>
                                        <div className="col-sm-2">
                                            <a className="element-box el-tablo centered trend-in-corner padded bold-label masariuman_colorYellow">
                                                <div className="value">
                                                    0
                                                </div>
                                                <div className="label">
                                                    Total Followers
                                                </div>
                                            </a>
                                        </div>
                                        <div className="col-sm-2">
                                            <a className="element-box el-tablo centered trend-in-corner padded bold-label masariuman_colorPurple">
                                                <div className="value">
                                                    0
                                                </div>
                                                <div className="label">
                                                    Total Genres
                                                </div>
                                            </a>
                                        </div>
                                        <div className="col-sm-3">
                                            <a className="element-box el-tablo centered trend-in-corner padded bold-label masariuman_colorOrange">
                                                <div className="value">
                                                    0
                                                </div>
                                                <div className="label">
                                                    Total Tags
                                                </div>
                                            </a>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div className="row">
                                <div className="col-sm-12">
                                    <div className="element-wrapper">
                                        <div className="col-sm-12">
                                            <a className="element-box el-tablo centered trend-in-corner padded bold-label masariuman_colorRed">
                                                <div className="value">
                                                    0
                                                </div>
                                                <div className="label">
                                                    Total Visitors
                                                </div>
                                            </a>
                                        </div>
                                    </div>
                                </div>
                            </div>
                            <div className="row">
                                <div className="col-sm-12">
                                    <div className="element-wrapper">
                                        {/* content here */}
                                        <div className="element-box">
                                            <h5 className="form-header">
                                                Translation's Dashboard
                                            </h5>
                                            <div className="form-desc">
                                                Manajemen Data Dashboard
                                            </div>
                                            <div className="os-tabs-w">
                                                <div className="os-tabs-controls">
                                                    <ul className="nav nav-tabs smaller">
                                                        <li className="nav-item">
                                                            <a className="nav-link active" data-toggle="tab">Visitor</a>
                                                        </li>
                                                    </ul>
                                                    <ul className="nav nav-pills smaller d-none d-md-flex">
                                                        <li className="nav-item">
                                                            <a className="nav-link active masariuman_cursorPointer" data-toggle="tab" >7 Days</a>
                                                        </li>
                                                        <li className="nav-item">
                                                            <a className="nav-link masariuman_cursorPointer" data-toggle="tab" >14 Days</a>
                                                        </li>
                                                        <li className="nav-item">
                                                            <a className="nav-link masariuman_cursorPointer" data-toggle="tab" >30 Days</a>
                                                        </li>
                                                    </ul>
                                                </div>
                                                <div className="tab-content">
                                                    <div className="tab-pane active" id="tab_overview">
                                                        <div className="el-tablo bigger">
                                                            <div className="label">
                                                                Translation Visitors
                                                            </div>
                                                            <div className="value">
                                                                0
                                                            </div>
                                                        </div>
                                                        <div className="el-chart-w">
                                                            <canvas height="150px" id="translationChart" width="600px"></canvas>
                                                        </div>
                                                    </div>
                                                    <div className="tab-pane" id="tab_sales"></div>
                                                    <div className="tab-pane" id="tab_conversion"></div>
                                                </div>
                                            </div>
                                        </div>
                                        {/* end content here */}
                                    </div>
                                </div>
                            </div>
                            <DarkMode />
                        </div>
                    </div>
                <Footer />
            </div>
        );
    }
}

export default LandingPage;
