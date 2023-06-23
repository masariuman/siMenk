import React, { Component } from "react";
import $ from 'jquery';
import Footer from "../../warudo/Footer";
import DarkMode from "../../warudo/DarkMode";
import { Link } from "react-router-dom";
import Loading from "../../warudo/Loading";
import swal from 'sweetalert2';
import Pagination from "react-js-pagination";
import axios from "axios";


class RumpunJabatan extends Component {
    constructor(props) {
        super(props);
        this.state = {
            rumpunJabatan: [],
            rumpun: "",
            rumpunJabatanEditInput: "",
            cari: "",
            totalItemsCount: 0,
            url: null,
            loading: false
        };
        this.handleChange = this.handleChange.bind(this);
        this.handleEditInputChange = this.handleEditInputChange.bind(this);
        this.handleEditSubmit = this.handleEditSubmit.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.renderRumpunJabatan = this.renderRumpunJabatan.bind(this);
        this.handlePageChange = this.handlePageChange.bind(this);
        this.modalTambah = this.modalTambah.bind(this);
        this.modalUbah = this.modalUbah.bind(this);
        this.handleChangeCari = this.handleChangeCari.bind(this);
    }

    handleChangeCari(e) {
        this.setState({
            cari: e.target.value
        });
        axios
            .post(`http://127.0.0.1:8877/v1/referensi/rumpun_jabatan/search`, {
                rumpun: e.target.value
            })
            .then(response => {
                // console.log(response.data);
                this.setState({
                    rumpunJabatan: response.data.data,
                    loading: false,
                    activePage: response.data.current_page,
                    itemsCountPerPage: response.data.per_page,
                    totalItemsCount: response.data.total,
                    pageRangeDisplayed: 10
                });
                // console.log(this.state.rumpunJabatan);
            });
    }

    handleDeleteButton(e) {
        // axios
        //     .get(`/masariuman_tag/${e}`)
        //     .then(response => {
        //         swal.fire({
        //             title: `Yakin ingin menghapus rumpunJabatan ${response.data.deeta_tag.rumpunJabatan}`,
        //             text: "Kalau Terhapus, Hubungi Admin Untuk Mengembalikan Data yang Terhapus!",
        //             icon: "warning",
        //             buttons: true,
        //             dangerMode: true,
        //           })
        //           .then((willDelete) => {
        //             if (willDelete) {
        //                 this.setState({
        //                     loading: true
        //                 });
        //                 axios
        //                     .delete(`/masariuman_tag/${e}`, {
        //                         url: this.state.url
        //                     })
        //                     .then(response => {
        //                         this.setState({
        //                             rumpunJabatan: response.data.deeta_tag.data,
        //                             loading: false
        //                         });
        //                         swal.fire("Sukses!", "Data Berhasil Dihapus!", "success");
        //                         // console.log("from handle sumit", response);
        //                     })
        //                     .catch(error => {
        //                         this.setState({
        //                             loading: false
        //                         });
        //                         swal.fire("Error!", "Gagal Menghapus Data, Silahkan Hubungi Admin!", "error");
        //                     });
        //             } else {
        //               swal.fire("Data Tidak Terhapus!");
        //             }
        //           });
        //     })
        //     .catch(error => {
        //         swal.fire("Error!", "Terdapat Masalah, Silahkan Hubungi Admin!", "error");
        //     });
    }

    handleEditButton(e) {
        axios
            .get(`http://127.0.0.1:8877/v1/referensi/rumpun_jabatan/${e}`)
            .then(response => {
                this.setState({
                    rumpunJabatanEditInput: response.data.data.rumpun,
                    rumpunId: response.data.data.id_rumpun
                });
            })
            .catch(error => {
                swal.fire("Error!", "Terdapat Masalah, Silahkan Hubungi Admin!", "error");
            });
    }

    handleChange(e) {
        this.setState({
            rumpun: e.target.value
        });
        // console.log(e.target.value);
    }

    handleEditInputChange(e) {
        this.setState({
            rumpunJabatanEditInput: e.target.value
        });
        // console.log(e.target.value);
    }

    handleSubmit(e) {
        e.preventDefault();
        this.setState({
            loading: true
        });
        axios
            .post("http://127.0.0.1:8877/v1/referensi/rumpun_jabatan", {
                rumpun: this.state.rumpun
            })
            .then(response => {
                this.setState({
                    rumpunJabatan: [response.data.data, ...this.state.rumpunJabatan],
                    rumpun: "",
                    loading: false
                });
                $("#tambahModal").removeClass("in");
                $(".modal-backdrop").remove();
                $('body').removeClass('modal-open');
                $('body').css('padding-right', '');
                $("#tambahModal").hide();
                swal.fire("Sukses!", "Data Baru Berhasil Ditambahkan!", "success");
                // console.log("from handle sumit", response);
            })
            .catch(error => {
                this.setState({
                    loading: false
                });
                var err = error.response.data.errors;
                var errStr = err.toString();
                swal.fire("Error!", errStr, "error");
                $("#tambahModal").removeClass("in");
                $(".modal-backdrop").remove();
                $('body').removeClass('modal-open');
                $('body').css('padding-right', '');
                $("#tambahModal").hide();
            });
        // console.log(this.state.rumpun);
    }

    handleEditSubmit(e) {
        e.preventDefault();
        this.setState({
            loading: true
        });
        axios
            .put(`http://127.0.0.1:8877/v1/referensi/rumpun_jabatan/${this.state.rumpunId}`, {
                rumpun: this.state.rumpunJabatanEditInput
            })
            .then(response => {
                this.setState({
                    rumpunJabatan: response.data.data,
                    rumpunJabatanEditInput: "",
                    loading: false
                });
                $("#editModal").removeClass("in");
                $(".modal-backdrop").remove();
                $('body').removeClass('modal-open');
                $('body').css('padding-right', '');
                $("#editModal").hide();
                swal.fire("Sukses!", "Data Berhasil Diubah!", "success");
                // console.log("from handle sumit", response);
            })
            .catch(error => {
                this.setState({
                    loading: false
                });
                swal.fire("Error!", "Gagal Mengubah Data, Silahkan Hubungi Admin!", "error");
            });
        // console.log(this.state.rumpun);
    }

    getRefRumpunJabatan() {
        this.setState({
            loading: true
        });
        axios
            .get("http://127.0.0.1:8877/v1/referensi/rumpun_jabatan")
            .then(response => {
                this.setState({
                    rumpunJabatan: response.data.data,
                    loading: false,
                    activePage: response.data.current_page,
                    itemsCountPerPage: response.data.per_page,
                    totalItemsCount: response.data.total,
                    pageRangeDisplayed: 10
                });
            })
            .catch(error => {
                swal.fire("Error!", "Terdapat Masalah, Silahkan Hubungi Admin!", "error");
                this.setState({loading: false});
            });
    }

    handlePageChange(pageNumber) {
        this.setState({
            loading: true
        });
        axios
            .get('http://127.0.0.1:8877/v1/referensi/rumpun_jabatan?page='+pageNumber)
            .then(response => {
                this.setState({
                    rumpunJabatan: response.data.data,
                    loading: false,
                    activePage: response.data.current_page,
                    itemsCountPerPage: response.data.per_page,
                    totalItemsCount: response.data.total,
                    pageRangeDisplayed: 10
                });
            })
            .catch(error => {
                swal.fire("Error!", "Terdapat Masalah, Silahkan Hubungi Admin!", "error");
                this.setState({loading: false});
            });
    }

    componentDidMount() {
        this.getRefRumpunJabatan();
        // console.log(this.state.rumpunJabatan);
    }

    componentDidUpdate() {
        // this.getRefRumpunJabatan();
    }

    renderRumpunJabatan() {
        return !this.state.rumpunJabatan ? <tr><td colSpan="3" className="text-center">Data Tidak Ditemukan</td></tr> :
            this.state.rumpunJabatan.map(rumpunJabatan => (
                <tr key={rumpunJabatan.id_rumpun}>
                    <th scope="row">{rumpunJabatan.nomor}</th>
                    <td>{rumpunJabatan.rumpun}</td>
                    <td>
                        <button data-target="#editModal" data-toggle="modal" className="mb-2 mr-2 border-0 btn-transition btn btn-shadow btn-outline-warning" type="button" onClick={this.handleEditButton.bind(this, rumpunJabatan.id_rumpun)}>Edit</button>
                        {/* <button className="mb-2 mr-2 border-0 btn-transition btn btn-shadow btn-outline-danger" type="button" onClick={this.handleDeleteButton.bind(this, rumpunJabatan.url)}>Delete</button> */}
                    </td>
                </tr>
            ));
    }

    modalTambah() {
        return (
            <div aria-hidden="true" className="onboarding-modal modal fade animated" id="tambahModal" role="dialog" tabIndex="-1">
                <div className="modal-dialog modal-lg modal-centered" role="document">
                    <div className="modal-content text-center">
                    <button aria-label="Close" className="close" data-dismiss="modal" type="button"><span className="close-label">Tutup</span><span className="os-icon os-icon-close"></span></button>
                    <div className="onboarding-side-by-side">
                        <div className="onboarding-media">
                        <img alt="" src="/iconModal/gridAdd.png" width="200px" />
                        </div>
                        <div className="onboarding-content with-gradient">
                        <h4 className="onboarding-title">
                            Tambah Rumpun Jabatan Baru
                        </h4>
                        <div className="onboarding-text">
                            Masukkan Rumpun Jabatan baru.
                        </div>
                        <form onSubmit={this.handleSubmit}>
                            <div className="row">
                            <div className="col-sm-12">
                                <div className="form-group">
                                    <input
                                        onChange={this.handleChange}
                                        value={this.state.rumpun}
                                        title="Rumpun Jabatan"
                                        placeholder="Masukkan Rumpun Jabatan Baru.."
                                        type="text"
                                        className="form-control"
                                    />
                                </div>
                            </div>
                            <div className="col-sm-12">
                                <div className="form-group text-center">
                                    <button className="mr-2 mb-2 btn btn-primary" data-target="#onboardingWideFormModal" data-toggle="modal" type="submit">Tambah  Rumpun Jabatan Baru</button>
                                </div>
                            </div>
                            </div>
                        </form>
                        </div>
                    </div>
                    </div>
                </div>
            </div>
        );
    }

    modalUbah() {
        return (
            <div aria-hidden="true" className="onboarding-modal modal fade animated" id="editModal" role="dialog" tabIndex="-1">
                <div className="modal-dialog modal-lg modal-centered" role="document">
                    <div className="modal-content text-center">
                    <button aria-label="Close" className="close" data-dismiss="modal" type="button"><span className="close-label">Tutup</span><span className="os-icon os-icon-close"></span></button>
                    <div className="onboarding-side-by-side">
                        <div className="onboarding-media">
                        <img alt="" src="/iconModal/gridAdd.png" width="200px" />
                        </div>
                        <div className="onboarding-content with-gradient">
                        <h4 className="onboarding-title">
                            Ubah Rumpun jabatan
                        </h4>
                        <div className="onboarding-text">
                            Masukkan Nama Baru Rumpun Jabatan.
                        </div>
                        <form onSubmit={this.handleEditSubmit}>
                            <div className="row">
                            <div className="col-sm-12">
                                <div className="form-group">
                                    <input
                                        onChange={this.handleEditInputChange}
                                        value={this.state.rumpunJabatanEditInput}
                                        title="Rumpun Jabatan"
                                        placeholder="Masukkan Nama Baru Rumpun Jabatan.."
                                        type="text"
                                        className="form-control"
                                    />
                                </div>
                            </div>
                            <div className="col-sm-12">
                                <div className="form-group text-center">
                                    <button className="mr-2 mb-2 btn btn-warning" data-target="#onboardingWideFormModal" data-toggle="modal" type="submit">Ubah Rumpun Jabatan</button>
                                </div>
                            </div>
                            </div>
                        </form>
                        </div>
                    </div>
                    </div>
                </div>
            </div>
        );
    }

    render() {
        return (
            this.state.loading === true ? <Loading /> :
            <div className="content-w">
                {/* title content */}
                <div className="top-bar color-scheme-transparent masariuman-height103px">
                    <div className="top-menu-controls masariuman-marginleft30px">
                        <div className="icon-w top-icon masariuman-titlecontent">
                        <div className="os-icon os-icon-rumpunJabatan"></div>
                        </div>
                        <div className="masariuman-textleft">
                            <span className="masariuman-bold"> RUMPUN JABATAN</span> <br/>
                            <small>Manajemen Rumpun Jabatan</small>
                        </div>
                    </div>
                    <div className="top-menu-controls">
                        <button className="mr-2 mb-2 btn btn-outline-primary" type="button" id="petunjuk"><i className="batch-icon-bulb-2"></i> PETUNJUK <i className="batch-icon-bulb"></i></button>
                    </div>
                </div>
                <ul className="breadcrumb">
                    <li className="breadcrumb-item">
                        <a> RUMPUN JABATAN</a>
                    </li>
                    <li className="breadcrumb-item">
                        <span>Manajemen Rumpun Jabatan</span>
                    </li>
                </ul>

                {/* content */}
                <div className="content-i masariuman-minheight100vh">
                        <div className="content-box">
                            <div className="element-wrapper">
                                {/* content here */}
                                <div className="element-box">
                                    <h5 className="form-header">
                                         Daftar Rumpun Jabatan
                                    </h5>
                                    <div className="form-desc">
                                        Data Manajemen Rumpun Jabatan
                                    </div>
                                    <div>
                                        <button className="mr-2 mb-2 btn btn-primary" data-target="#tambahModal" data-toggle="modal" type="button" id="buttonTambahModal">Tambah  Rumpun Jabatan Baru</button>
                                        <div className="col-sm-4 float-right">
                                            <input type="text" className="form-control" onChange={this.handleChangeCari}
                                                value={this.state.cari} placeholder="Cari  Rumpun Jabatan..."></input>
                                        </div>
                                    </div>
                                    <div className="table-responsive" id="ruanganTable">
                                        <table id="tabeldata" width="100%" className="table table-striped table-lightfont">
                                            <thead>
                                                <tr>
                                                    <th className="width50px">NO</th>
                                                    <th> NAMA RUMPUN JABATAN</th>
                                                    <th className="width250px">AKSI</th>
                                                </tr>
                                            </thead>
                                            <tbody>{this.renderRumpunJabatan()}</tbody>
                                        </table>
                                    </div>
                                    <div className="d-flex justify-content-center">
                                        <Pagination
                                            activePage={this.state.activePage}
                                            itemsCountPerPage={this.state.itemsCountPerPage}
                                            totalItemsCount={this.state.totalItemsCount}
                                            pageRangeDisplayed={this.state.pageRangeDisplayed}
                                            onChange={this.handlePageChange}
                                        />
                                    </div>
                                </div>
                                {/* end content here */}
                            </div>
                            <DarkMode />
                        </div>
                    </div>
                <Footer />
                {this.modalTambah()}
                {this.modalUbah()}
            </div>
        );
    }
}

export default RumpunJabatan;
