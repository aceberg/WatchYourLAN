var histArray = {};

loadHistory();

function createHTML(hist, i) {
    
    let allState = "";
    let color = "";

    for (let state of hist.State){
        if (state.State) {
            color = `bi-check-circle-fill" style="color:var(--bs-success);"`;
        } else {
            color = `bi-dash-circle-fill" style="color:var(--bs-danger);"`;
        }
        allState = allState + `<i class="bi ${color} title="${state.Date}"></i>`;
    }

    let html = `
    <tr>
        <td style="opacity: 45%;">${i}.</td>
        <td><a href="/scan/?addr=${hist.Addr}">${hist.Name}</a></td>
        <td><a href="/scan/?addr=${hist.Addr}">${hist.Addr}</a></td>
        <td><a href="http://${hist.Addr}:${hist.Port}">${hist.Port}</a></td>
        <td><a href="http://${hist.Addr}:${hist.Port}">${hist.PortName}</a></td>
        <td>${allState}</td>
    </tr>`;
    
    return html;
}

async function loadHistory() {
    
    let url = '/api/history';
    let histMap = await (await fetch(url)).json();
    if (histMap != null) {
        histArray = Object.values(histMap);
    }

    displayArrayData(histArray);
}

function sortBy(field) {
    sortByAny(histArray, field);
}