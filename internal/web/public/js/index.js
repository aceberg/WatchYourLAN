var addrsArray = {};

loadAddrs();

function createHTML(addr, i) {
    let now = '';

    if (addr.Now == 0) {
        now = `<i class="bi bi-circle-fill" style="color:var(--bs-gray-500);"></i>`;
    } else {
        now = `<i class="bi bi-check-circle-fill" style="color:var(--bs-success);"></i>`;
    }
    
    let html = `
    <tr>
        <td style="opacity: 45%;">${i}.</td>
        <td>${addr.Name}</td>
        <td>
            <a href="http://${addr.IP}">${addr.IP}</a>
        </td>
        <td>${addr.Mac}</td>
        <td>${addr.Hw}</td>
        <td>${addr.Date}</td>
        <td>${addr.Known}</td>
        <td>${now}</td>
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