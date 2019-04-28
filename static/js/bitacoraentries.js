$(document).ready(function () {

    $('#alert').hide();
    $('#alert_error').hide();

    if (existsEmployeeParamInURI()) {
        console.log('Exists employee query string ... ');
        const urlParams = new URLSearchParams(window.location.search);
        var empID = urlParams.get('employee');
        var url = 'http://localhost:8083/api/entries/' + empID; 
        populateEntriesTable(url, true);
    } else {
        console.log('It does not exists ...');
        var url = 'http://localhost:8083/api/entries/';
        populateEntriesTable(url, false);
    }

});

function populateEntriesTable(url, single) {
    $.ajax({
        url: url,
        type: 'GET',
        data: {},
        success: function(data) {
            var entriesTable = $('#entries_table');
            if (single) {
                showInfoInTable(entriesTable, data.entries, data);
            } else {
                showInfoInTableForAll(entriesTable, data);
            }
        },
        error: function(data) {
            console.log('woops! :(');
            console.log(data);
        }
    });
}

function showInfoInTable(entriesTable, entries, data) {
    for (var i = 0; i < entries.length; i++) {
        var entry = entries[i];
        var whoRow = '<tr><td><b><a href="/export?id=' + data._id + '">' + data.name + '</a></b></td>';
        whoRow += '<td>' + entry.description + '</td>';
        whoRow += '<td>' + entry.date + '</td></tr>';
        entriesTable.append(whoRow);
    }
}

function showInfoInTableForAll(entriesTable, data) {
    for (var i = 0; i < data.length; i++) {
        showInfoInTable(entriesTable, data[i].entries, data[i]);
    }
}

function existsEmployeeParamInURI() {
    var url = new URL(window.location.href);
    if (url.searchParams.get('employee')) {
        return true;
    } else {
        return false;
    }
}

