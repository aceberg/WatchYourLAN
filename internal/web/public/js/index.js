var addrsArray = {};

loadAddrs();

function createHTML(addr, i) {
    let off = '';

    if (addr.Offline == 0) {
        off = `<span style="color:var(--bs-success);">0</span>`;
    } else {
        off = `<span style="color:var(--bs-danger);">${addr.Offline}</span>`;
    }
    
    let html = `
    <tr>
        <td style="opacity: 45%;">${i}.</td>
        <td>
            <a href="/scan/?addr=${addr.Addr}">${addr.Name}</a>
        </td>
        <td>
            <a href="/scan/?addr=${addr.Addr}">${addr.Addr}</a>
        </td>
        <td>${addr.Total}</td>
        <td>${addr.Watching}</td>
        <td>${addr.Online}</td>
        <td>${off}</td>
    </tr>
    `;
    
    return html;
}

async function loadAddrs() {
    
    let url = '/api/all';
    let addrsMap = await (await fetch(url)).json();
    if (addrsMap != null) {
        addrsArray = Object.values(addrsMap);
    }

    displayArrayData(addrsArray);
}

function sortBy(field) {
    sortByAny(addrsArray, field);
}