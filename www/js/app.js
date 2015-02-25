$(document).ready(function() {
    var winnerTableRow = winnerTableRowTemplate();
    var looserTableRow = looserTableRowTemplate();

    $.get("users", function(users) {
        $.get("scores", function(scores) {
            users = JSON.parse(users);
            scores = JSON.parse(scores);

            var usersById = normalizeUsers(users);

            updateRankingTables(scores, usersById);
        });
    });

    function normalizeUsers(users) {
        var normalized = new Array;

        for (i in users) {
            normalized[users[i].Id] = users[i];
        }

        return normalized;
    }

    function composeData(user, score) {
        return {
            place: parseInt(i) + 1,
            avatar_url: "http://zadane.pl/img/avatars/100-ONA.png",
            nick: user['Nick'],
            name: user['Name'],
            answers7days: score['Answers7Days'],
            answers30days: score['Answers30Days'],
            answers90days: score['Answers90Days']
        };
    }

    function updateRankingTables(scores, users) {
        var score, user, data;

        for (i in scores) {
            score = scores[i];
            user = users[score["UserId"]];
            data = {};

            data = composeData(user, score);
            appendTableRowForPlace(i, data);
        }
    }

    function appendTableRowForPlace(place, data) {
        if (place < 3) {
            var row = winnerTableRow(data);
            $(".winners-table > tbody:last").append(row);
        } else {
            var row = looserTableRow(data);
            $(".loosers-table > tbody:last").append(row);
        }
    }

    function winnerTableRowTemplate() {
        return _.template('<tr>' +
                   '<td class="place"><%= place %></td>' +
                   '<td class="avatar">' +
                       '<div class="avatar-image">' +
                           '<img src="<%= avatar_url %>" class="img-circle avatar-img" />' +
                           '<img src="img/icon_prize_0<%= place %>.png" class="prize-img" />' +
                       '</div>' +
                       '<div class="user-name">' +
                           '<p class="nick"><%= nick %></p>' +
                           '<p class="name"><%= name %></p>' +
                       '</div>' +
                   '</td>' +
                   '<td class="score"><%= answers7days %></td>' +
                   '<td class="score"><%= answers30days %></td>' +
                   '<td class="score"><%= answers90days %></td>' +
               '</tr>');
    }

    function looserTableRowTemplate() {
        return _.template('<tr>' +
                    '<td class="place"><%= place %></td>' +
                    '<td class="avatar">' +
                        '<img src="<%= avatar_url %>" class="img-circle" />' +
                        '<div class="user-name">' +
                            '<p class="nick"><%= nick %></p>' +
                            '<p class="name"><%= name %></p>' +
                        '</div>' +
                    '</td>' +
                    '<td class="score"><%= answers7days %></td>' +
                    '<td class="score"><%= answers30days %></td>' +
                    '<td class="score"><%= answers90days %></td>' +
                '</tr>');
    }
});
