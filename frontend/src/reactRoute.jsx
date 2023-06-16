import React, { Component } from "react";
import {
    BrowserRouter as Router,
    Routes,
    Route
} from "react-router-dom";
import MobileMenu from "./warudo/MobileMenu";
import Menu from "./warudo/Menu";
import Empatkosongempat from "./warudo/Empatkosongempat";
import LandingPage from "./component/LandingPage";
import ReferensiJafung from "./component/referensi/Jafung"
import ReferensiRumpunJabatan from "./component/referensi/RumpunJabatan"
// import DashboardIndex from "./component/dashboard/Index";

class ReactRoute extends Component {
    constructor(props) {
        super(props);
        this.state = {
            level: ""
        };
    }

    // getUuzaa() {
    //     axios.get("/getUuzaa").then((response) => {
    //         this.setState({
    //             level: response.data.data.reberu
    //         });
    //     });
    // }

    componentDidMount() {
        // this.getUuzaa();
    }

    render() {
        return (
            <Router>
                <MobileMenu />
                <Menu />
                <Routes>
                    <Route exact path="/" element={< LandingPage />} />
                    <Route exact path="/referensi/rumpun_jabatan" element={< ReferensiRumpunJabatan />} />
                    <Route exact path="/referensi/jafung" element={< ReferensiJafung />} />
                    <Route path="*" element={< Empatkosongempat />}/>
                </Routes>
            </Router>
            );
    }
}

export default ReactRoute;
