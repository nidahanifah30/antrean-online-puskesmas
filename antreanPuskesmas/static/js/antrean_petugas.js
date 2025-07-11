$(document).ready(function () {
  // Isi data ke modal edit
$('#modalEdit').on('show.bs.modal', function (event) {
    const button = $(event.relatedTarget)
    const id = button.data('id')
    const nama = button.data('nama')
    const layanan = button.data('layanan')
    const nik = button.data('nik')
    const nohp = button.data('nohp')
    const tanggal = button.data('tanggal')

    $('#edit-id').val(id)
    $('#edit-nama').val(nama)
    $('#edit-layanan').val(layanan)
    $('#edit-nik').val(nik)
    $('#edit-nohp').val(nohp)
    $('#edit-tanggal').val(tanggal)
})

    // Modal konfirmasi (hapus, panggil, selesai)
    ['Hapus', 'Panggil', 'Selesai'].forEach(function (type) {
        const modalId = '#modal' + type
        $(modalId).on('show.bs.modal', function (event) {
            const button = $(event.relatedTarget)
            const id = button.data('id')
            const nama = button.data('nama')

            $(modalId + ' .modal-body span').text(nama)
            $(modalId + ' form').attr('action', '/petugas/' + type.toLowerCase() + '/' + id)
        })
    })
})

$('#modalEdit').on('show.bs.modal', function (event) {
    const button = $(event.relatedTarget)
    const id = button.data('id')

    $('#edit-id').val(id)
    $('#edit-nama').val(button.data('nama'))
    $('#edit-nik').val(button.data('nik'))
    $('#edit-nohp').val(button.data('nohp'))
    $('#edit-layanan').val(button.data('layanan'))
    $('#edit-tanggal').val(button.data('tanggal'))

    $('#formEdit').attr('action', '/petugas/update/' + id)
})

$('#modalHapus').on('show.bs.modal', function (event) {
    const button = $(event.relatedTarget)
    const id = button.data('id')
    const nama = button.data('nama')

    $('#hapus-nama').text(nama)
    $('#formHapus').attr('action', '/petugas/delete/' + id)
})

$('#modalPanggil').on('show.bs.modal', function (event) {
    const button = $(event.relatedTarget)
    const id = button.data('id')
    const nama = button.data('nama')

    $('#panggil-nama').text(nama)
    $('#formPanggil').attr('action', '/petugas/panggil/' + id)
})

$('#modalSelesai').on('show.bs.modal', function (event) {
    const button = $(event.relatedTarget)
    const id = button.data('id')
    const nama = button.data('nama')

    $('#selesai-nama').text(nama)
    $('#formSelesai').attr('action', '/petugas/selesai/' + id)
})
