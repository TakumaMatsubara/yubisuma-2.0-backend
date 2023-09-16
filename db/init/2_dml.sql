USE `yubisuma_api`;

SET NAMES utf8mb4;

INSERT INTO
    `animals` (
        `id`,
        `name`,
        `numFinger`,
        `skill`,
        `skillDesc`,
        `handUrl`,
        `upFingerUrl`,
        `downFingerUrl`
    )
VALUES (
        UUID(),
        "ゴリラ",
        5,
        "強力な握力",
        "ゴリラの握力は非常に強く、大きな物も簡単に掴むことができます。",
        "handUrl_for_ゴリラ",
        "upFingerUrl_for_ゴリラ",
        "downFingerUrl_for_ゴリラ"
    ), (
        UUID(),
        "いか",
        10,
        "インク吹き",
        "いかは危険を感じるとインクを吹いて敵を混乱させます。",
        "handUrl_for_いか",
        "upFingerUrl_for_いか",
        "downFingerUrl_for_いか"
    ), (
        UUID(),
        "たこ",
        8,
        "吸盤",
        "たこの足には吸盤があり、これによって様々な物に吸い付くことができます。",
        "handUrl_for_たこ",
        "upFingerUrl_for_たこ",
        "downFingerUrl_for_たこ"
    ), (
        UUID(),
        "ライオン",
        5,
        "強烈な咆哮",
        "ライオンの咆哮は非常に強烈で、敵を怯ませる効果があります。",
        "handUrl_for_ライオン",
        "upFingerUrl_for_ライオン",
        "downFingerUrl_for_ライオン"
    ), (
        UUID(),
        "人",
        5,
        "知恵",
        "人は知恵を持っており、様々な問題を解決する能力があります。",
        "handUrl_for_人",
        "upFingerUrl_for_人",
        "downFingerUrl_for_人"
    ), (
        UUID(),
        "へび",
        0,
        "毒",
        "一部のへびには強力な毒を持っており、これによって獲物を捕らえます。",
        "handUrl_for_へび",
        "upFingerUrl_for_へび",
        "downFingerUrl_for_へび"
    );