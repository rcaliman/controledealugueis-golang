function exibeComboMeses() {
    meses = ['janeiro', 'fevereiro', 'março', 'abril', 'maio', 'junho', 'julho', 'agosto', 'setembro', 'outubro', 'novembro', 'dezembro']
    let data = new Date()
    mesAtual = data.getMonth()
    let html = "<select id='selectMeses' name='mes' class='form-control btn-navegate'>" + "<option value='" + meses[mesAtual] + "'>" + meses[mesAtual] + "</option>"
    for (let i = 0; i < meses.length; i++) {
        html += "<option value='" + meses[i] + "'>" + meses[i] + "</option>"
    }
    html += "</select>"
    document.getElementById("comboMeses").innerHTML = html
}

function exibeComboAnos() {
    let data = new Date()
    anoAtual = data.getFullYear()
    let html = "<select id='selectAnos' name='ano' class='form-control btn-navegate'>"
    html += "<option value='" + anoAtual + "'>" + anoAtual + "</option>"
    html += "<option value='" + (anoAtual - 1) + "'>" + (anoAtual - 1) + "</option>"
    html += "<option value='" + anoAtual + "'>" + anoAtual + "</option>"
    html += "<option value='" + (anoAtual + 1) + "'>" + (anoAtual + 1) + "</option>"
    html += "</select>"
    document.getElementById("comboAnos").innerHTML = html
}