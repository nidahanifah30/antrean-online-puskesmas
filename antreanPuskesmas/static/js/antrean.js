let antreanData = {}

$('#form-antrean').on('submit', function (e) {
    e.preventDefault()

    const formData = $(this).serialize()

    $.post('/antrean/preview', formData, function (res) {
        antreanData = res

        $('#popup-nama').text(res.nama)
        $('#popup-layanan').text(res.layanan)
        $('#popup-tanggal').text(res.tanggal)
        $('#popup-nomor').text(res.nomor)

        const modal = new bootstrap.Modal(document.getElementById('modalAntrean'))
        modal.show()
    })
    })

    $('#btn-simpan').on('click', function () {
    $.post('/antrean/simpan', antreanData, function () {
        alert('Antrean berhasil disimpan!')
        window.location.href = '/dashboard'
    })
})

$('#btn-refresh').on('click', function () {
    $.get('/antrean/status/data', function (data) {
        $('#status-antrean').text(data.status)
        $('#nomor-antrean').text(data.nomor)
        $('#estimasi-waktu').text(data.estimasiMenit + ' menit')
    }).fail(function () {
        alert('Gagal memuat data status.')
    })
})



