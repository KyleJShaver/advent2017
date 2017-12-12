package day12

import (
	"fmt"
	"testing"
)

type Example struct {
	input string
	part1 int
	part2 int
}

var personalData = "0 <-> 659, 737\n1 <-> 1, 1433\n2 <-> 982, 1869\n3 <-> 306, 380, 1462, 1827\n4 <-> 1076\n5 <-> 794, 1451\n6 <-> 146, 1055\n7 <-> 834, 1557\n8 <-> 1333\n9 <-> 849, 906, 1863\n10 <-> 362, 505\n11 <-> 33, 938, 1896\n12 <-> 490, 913\n13 <-> 189, 690\n14 <-> 796\n15 <-> 56, 280, 1288, 1721\n16 <-> 16\n17 <-> 904\n18 <-> 150, 1394, 1458\n19 <-> 1773\n20 <-> 70\n21 <-> 993\n22 <-> 22\n23 <-> 285, 1004\n24 <-> 209, 727\n25 <-> 614, 1590\n26 <-> 113\n27 <-> 1321, 1341\n28 <-> 351, 730, 1037\n29 <-> 29, 1828\n30 <-> 1378, 1983\n31 <-> 705, 1035\n32 <-> 1042, 1519\n33 <-> 11, 886\n34 <-> 360, 1101, 1531, 1877\n35 <-> 971, 1652\n36 <-> 188\n37 <-> 1935\n38 <-> 38\n39 <-> 39, 1472\n40 <-> 84, 1110\n41 <-> 483, 895, 1035, 1864, 1919\n42 <-> 624, 813, 1484, 1517\n43 <-> 492\n44 <-> 947, 1572\n45 <-> 45, 1589, 1748, 1836\n46 <-> 1821\n47 <-> 123, 206, 239\n48 <-> 146\n49 <-> 235, 871\n50 <-> 172, 1672\n51 <-> 504\n52 <-> 585, 678, 878\n53 <-> 484, 543, 1282\n54 <-> 374, 723\n55 <-> 1956\n56 <-> 15\n57 <-> 583, 1159, 1596\n58 <-> 313, 580, 1101\n59 <-> 1529, 1788\n60 <-> 60\n61 <-> 1033, 1857\n62 <-> 874\n63 <-> 63, 1007, 1316, 1673\n64 <-> 1140\n65 <-> 1136\n66 <-> 1337, 1546\n67 <-> 67, 1468\n68 <-> 298, 951\n69 <-> 906\n70 <-> 20, 1977\n71 <-> 1956\n72 <-> 146, 1465\n73 <-> 911\n74 <-> 1345\n75 <-> 711\n76 <-> 732, 789, 1499, 1637\n77 <-> 1897\n78 <-> 78, 379\n79 <-> 96, 462, 847\n80 <-> 886, 907\n81 <-> 1564\n82 <-> 1362, 1680\n83 <-> 225, 916\n84 <-> 40, 1460\n85 <-> 330\n86 <-> 1370\n87 <-> 87, 640\n88 <-> 806, 1411\n89 <-> 1732\n90 <-> 603\n91 <-> 547, 904\n92 <-> 1106\n93 <-> 782\n94 <-> 401, 1039, 1148, 1356\n95 <-> 95, 344, 1092\n96 <-> 79\n97 <-> 1555, 1649\n98 <-> 133\n99 <-> 866\n100 <-> 116, 1183, 1450\n101 <-> 964\n102 <-> 521, 843\n103 <-> 1002\n104 <-> 675, 1752, 1760\n105 <-> 447\n106 <-> 625\n107 <-> 200, 923, 1573\n108 <-> 108, 1535\n109 <-> 1938, 1984\n110 <-> 992\n111 <-> 410, 436, 1789, 1985\n112 <-> 730, 1742\n113 <-> 26, 774, 1620\n114 <-> 891\n115 <-> 556, 980, 1502\n116 <-> 100, 405, 438\n117 <-> 1220, 1747\n118 <-> 417\n119 <-> 1671\n120 <-> 225\n121 <-> 215, 1414\n122 <-> 1421\n123 <-> 47, 488, 1390\n124 <-> 1750\n125 <-> 1020, 1456, 1645, 1811\n126 <-> 945, 988, 1558\n127 <-> 1562\n128 <-> 246, 419, 878, 1057\n129 <-> 198, 214, 526\n130 <-> 1572\n131 <-> 624\n132 <-> 1086\n133 <-> 98, 430, 1950\n134 <-> 1179\n135 <-> 1603\n136 <-> 136\n137 <-> 137, 248\n138 <-> 604\n139 <-> 1080, 1744, 1829\n140 <-> 786, 890, 1336\n141 <-> 819, 835\n142 <-> 142, 1657\n143 <-> 590, 670\n144 <-> 622\n145 <-> 182, 930, 1164, 1741\n146 <-> 6, 48, 72, 152\n147 <-> 1880\n148 <-> 1486\n149 <-> 1595\n150 <-> 18, 653\n151 <-> 1834\n152 <-> 146, 289, 1949\n153 <-> 1379\n154 <-> 694, 1025\n155 <-> 1143, 1469\n156 <-> 437, 1492, 1616\n157 <-> 1044\n158 <-> 410, 1391\n159 <-> 1327\n160 <-> 387, 892, 963, 1287\n161 <-> 964, 1017\n162 <-> 786, 1098, 1351, 1445, 1508\n163 <-> 415\n164 <-> 255, 790, 1410\n165 <-> 252, 425, 1186, 1662, 1838\n166 <-> 791, 1012\n167 <-> 167, 836, 1922\n168 <-> 1586, 1998\n169 <-> 679, 914\n170 <-> 1975\n171 <-> 672\n172 <-> 50\n173 <-> 614\n174 <-> 723\n175 <-> 246, 747\n176 <-> 400, 1338\n177 <-> 573, 1617, 1724\n178 <-> 473, 1572\n179 <-> 183\n180 <-> 200, 1380\n181 <-> 1394, 1657\n182 <-> 145, 1825\n183 <-> 179, 399, 955, 1236, 1295, 1840\n184 <-> 712\n185 <-> 185\n186 <-> 551, 1627\n187 <-> 1195\n188 <-> 36, 440, 1277, 1362\n189 <-> 13\n190 <-> 1111\n191 <-> 985, 1372\n192 <-> 496, 1137, 1283\n193 <-> 827, 1053\n194 <-> 610\n195 <-> 1231, 1497\n196 <-> 1960\n197 <-> 584, 1359\n198 <-> 129, 651, 714, 1663\n199 <-> 1047, 1805\n200 <-> 107, 180, 658, 1397\n201 <-> 1590\n202 <-> 1268, 1768\n203 <-> 1683\n204 <-> 567, 1848\n205 <-> 276\n206 <-> 47, 461, 1794\n207 <-> 207\n208 <-> 1248\n209 <-> 24, 1589\n210 <-> 1063, 1504\n211 <-> 907, 1815\n212 <-> 1948\n213 <-> 388\n214 <-> 129, 561, 793, 1569\n215 <-> 121, 252\n216 <-> 216, 1728\n217 <-> 761\n218 <-> 631, 816\n219 <-> 459, 807, 1008\n220 <-> 1231, 1447\n221 <-> 713, 1541\n222 <-> 856, 924, 1924\n223 <-> 1103\n224 <-> 1426, 1761\n225 <-> 83, 120, 242, 784\n226 <-> 311, 560\n227 <-> 587, 667, 984, 1091\n228 <-> 1852\n229 <-> 229, 1198, 1204\n230 <-> 1505, 1944\n231 <-> 1158, 1594, 1925\n232 <-> 232, 345, 1417\n233 <-> 828, 1677\n234 <-> 982\n235 <-> 49, 1012, 1254, 1956\n236 <-> 489, 893, 1545\n237 <-> 988, 1450\n238 <-> 1719, 1791\n239 <-> 47, 736, 1027\n240 <-> 1167, 1457\n241 <-> 270\n242 <-> 225, 814, 1873\n243 <-> 243, 282, 914\n244 <-> 302, 328, 559\n245 <-> 1753\n246 <-> 128, 175\n247 <-> 1230, 1660, 1824\n248 <-> 137, 537, 1423\n249 <-> 1209, 1391, 1749\n250 <-> 298\n251 <-> 799\n252 <-> 165, 215\n253 <-> 1437, 1741\n254 <-> 1356\n255 <-> 164\n256 <-> 310\n257 <-> 1267, 1821\n258 <-> 1177, 1876\n259 <-> 1257, 1432\n260 <-> 311, 1115\n261 <-> 1504\n262 <-> 1268\n263 <-> 565, 1023\n264 <-> 408, 1806\n265 <-> 1009, 1144\n266 <-> 599, 616\n267 <-> 1182\n268 <-> 1026, 1456, 1470, 1854\n269 <-> 269, 686\n270 <-> 241, 1445\n271 <-> 1232\n272 <-> 487, 863, 1286\n273 <-> 1614, 1748\n274 <-> 1332\n275 <-> 1010, 1334\n276 <-> 205, 478, 888, 1049\n277 <-> 277, 1330\n278 <-> 734, 1561\n279 <-> 535, 1190, 1913\n280 <-> 15\n281 <-> 1206\n282 <-> 243, 475, 1571\n283 <-> 283, 1936\n284 <-> 324, 702, 844, 1601\n285 <-> 23\n286 <-> 286, 1080, 1127\n287 <-> 295\n288 <-> 805\n289 <-> 152, 656, 691, 993\n290 <-> 595, 616, 1199\n291 <-> 1028, 1218, 1844\n292 <-> 1447\n293 <-> 378, 1771\n294 <-> 811, 1790\n295 <-> 287, 367, 693\n296 <-> 861, 1948, 1950\n297 <-> 1631\n298 <-> 68, 250, 641, 745\n299 <-> 299, 898, 1152, 1574\n300 <-> 353\n301 <-> 990, 1340, 1960\n302 <-> 244\n303 <-> 1645\n304 <-> 445, 448, 557, 1611\n305 <-> 1350, 1442\n306 <-> 3, 445\n307 <-> 567\n308 <-> 1799\n309 <-> 864\n310 <-> 256, 853, 1911\n311 <-> 226, 260, 970\n312 <-> 1308\n313 <-> 58\n314 <-> 1795\n315 <-> 1225, 1627, 1903\n316 <-> 1065, 1991\n317 <-> 317\n318 <-> 546, 1415\n319 <-> 1883\n320 <-> 417, 1040\n321 <-> 1927\n322 <-> 686\n323 <-> 1018, 1619\n324 <-> 284\n325 <-> 1114\n326 <-> 474, 872\n327 <-> 439\n328 <-> 244\n329 <-> 329\n330 <-> 85, 1868\n331 <-> 331, 680, 1224, 1243, 1291\n332 <-> 1748\n333 <-> 1673\n334 <-> 676, 1306\n335 <-> 809\n336 <-> 797, 1559, 1858\n337 <-> 978, 1874\n338 <-> 538, 1077\n339 <-> 432, 1033\n340 <-> 340\n341 <-> 1267, 1683\n342 <-> 587, 726\n343 <-> 1560, 1705\n344 <-> 95\n345 <-> 232, 1032\n346 <-> 581, 1992\n347 <-> 530\n348 <-> 667, 1479\n349 <-> 1620\n350 <-> 861, 998, 1069\n351 <-> 28\n352 <-> 352\n353 <-> 300, 1513\n354 <-> 1109, 1747\n355 <-> 950, 965, 1394, 1616\n356 <-> 1136\n357 <-> 1443, 1698\n358 <-> 973, 1814\n359 <-> 1662\n360 <-> 34, 921\n361 <-> 1418\n362 <-> 10, 668, 1656\n363 <-> 450, 766\n364 <-> 615, 1941\n365 <-> 1071\n366 <-> 749, 1375\n367 <-> 295, 1708, 1925\n368 <-> 383, 928\n369 <-> 369\n370 <-> 651, 1048\n371 <-> 665, 1460, 1686\n372 <-> 390, 958, 1682\n373 <-> 1559\n374 <-> 54, 1395, 1847\n375 <-> 799, 1061, 1383, 1773\n376 <-> 376\n377 <-> 1463\n378 <-> 293, 872, 917\n379 <-> 78, 449\n380 <-> 3, 1117\n381 <-> 1995\n382 <-> 846, 896, 1506, 1918\n383 <-> 368, 477\n384 <-> 1009, 1788\n385 <-> 1926\n386 <-> 1602\n387 <-> 160, 667\n388 <-> 213, 713\n389 <-> 586\n390 <-> 372, 1140\n391 <-> 863\n392 <-> 590, 1259, 1384\n393 <-> 645\n394 <-> 423, 1537\n395 <-> 782, 1496, 1893\n396 <-> 1497\n397 <-> 397\n398 <-> 847, 1265\n399 <-> 183, 743, 1564\n400 <-> 176\n401 <-> 94\n402 <-> 402\n403 <-> 558\n404 <-> 572\n405 <-> 116\n406 <-> 769, 1563\n407 <-> 1790\n408 <-> 264\n409 <-> 847\n410 <-> 111, 158, 1467, 1880, 1902\n411 <-> 1221\n412 <-> 858, 1088, 1848\n413 <-> 1589\n414 <-> 468, 637\n415 <-> 163, 415\n416 <-> 416, 817\n417 <-> 118, 320, 1671\n418 <-> 1029, 1987\n419 <-> 128, 765\n420 <-> 420, 1180\n421 <-> 1170, 1766\n422 <-> 563, 1400, 1904, 1926\n423 <-> 394\n424 <-> 730\n425 <-> 165, 478, 957\n426 <-> 1505\n427 <-> 1478, 1780\n428 <-> 1658\n429 <-> 650\n430 <-> 133\n431 <-> 469, 1491\n432 <-> 339\n433 <-> 667, 1549\n434 <-> 852\n435 <-> 435\n436 <-> 111\n437 <-> 156\n438 <-> 116, 874\n439 <-> 327, 1387\n440 <-> 188, 540\n441 <-> 441\n442 <-> 692\n443 <-> 1037, 1153\n444 <-> 758, 1064\n445 <-> 304, 306\n446 <-> 1072, 1495, 1890\n447 <-> 105, 1665\n448 <-> 304, 591\n449 <-> 379\n450 <-> 363, 450, 466\n451 <-> 1135\n452 <-> 990, 1344, 1604\n453 <-> 1390, 1755\n454 <-> 1853\n455 <-> 1199, 1735, 1852\n456 <-> 484\n457 <-> 457, 1973\n458 <-> 786\n459 <-> 219\n460 <-> 1655, 1777\n461 <-> 206, 753\n462 <-> 79, 668\n463 <-> 699\n464 <-> 1843\n465 <-> 829, 978, 1820\n466 <-> 450\n467 <-> 770, 1475\n468 <-> 414, 613, 1230\n469 <-> 431, 1470\n470 <-> 1292, 1659\n471 <-> 603\n472 <-> 719, 1398\n473 <-> 178, 1419\n474 <-> 326, 866\n475 <-> 282\n476 <-> 1512, 1669\n477 <-> 383\n478 <-> 276, 425\n479 <-> 1029\n480 <-> 603\n481 <-> 991, 1878\n482 <-> 1726, 1783\n483 <-> 41, 855\n484 <-> 53, 456\n485 <-> 1814\n486 <-> 1470\n487 <-> 272, 1693\n488 <-> 123\n489 <-> 236, 1155, 1793\n490 <-> 12, 1349, 1498\n491 <-> 913\n492 <-> 43, 1121, 1822\n493 <-> 1148\n494 <-> 619, 1528\n495 <-> 1960\n496 <-> 192\n497 <-> 497\n498 <-> 1185\n499 <-> 1933\n500 <-> 1948\n501 <-> 1364, 1433, 1826\n502 <-> 893, 933\n503 <-> 514\n504 <-> 51, 821, 1116\n505 <-> 10, 905\n506 <-> 506\n507 <-> 892\n508 <-> 508\n509 <-> 1078\n510 <-> 1716\n511 <-> 1197, 1900\n512 <-> 568, 1284, 1422\n513 <-> 1604, 1817, 1927\n514 <-> 503, 773, 1725, 1883\n515 <-> 1396\n516 <-> 542\n517 <-> 1299\n518 <-> 1562\n519 <-> 1622\n520 <-> 1254, 1868, 1881\n521 <-> 102, 1765\n522 <-> 1438, 1980\n523 <-> 1779\n524 <-> 1235, 1968\n525 <-> 1267, 1484\n526 <-> 129\n527 <-> 1211\n528 <-> 597, 1017\n529 <-> 529, 609\n530 <-> 347, 1737, 1890\n531 <-> 851, 1320\n532 <-> 709\n533 <-> 819, 1091\n534 <-> 856\n535 <-> 279, 535\n536 <-> 536\n537 <-> 248, 1317, 1763\n538 <-> 338, 1793\n539 <-> 539, 1485\n540 <-> 440, 1216\n541 <-> 541\n542 <-> 516, 1232, 1831\n543 <-> 53, 1234\n544 <-> 639, 1090\n545 <-> 1898\n546 <-> 318, 1208\n547 <-> 91\n548 <-> 569, 750, 1090\n549 <-> 1720\n550 <-> 660\n551 <-> 186\n552 <-> 642, 672, 723\n553 <-> 1927\n554 <-> 717, 1841, 1997\n555 <-> 1938\n556 <-> 115, 1453\n557 <-> 304, 904\n558 <-> 403, 1438\n559 <-> 244, 1046\n560 <-> 226\n561 <-> 214, 793, 1570\n562 <-> 1992\n563 <-> 422\n564 <-> 860\n565 <-> 263, 804\n566 <-> 566\n567 <-> 204, 307\n568 <-> 512, 1084\n569 <-> 548, 985, 1910\n570 <-> 665, 1844\n571 <-> 735\n572 <-> 404, 1026, 1111\n573 <-> 177, 1994\n574 <-> 846\n575 <-> 723, 738, 795, 802, 926, 1215\n576 <-> 1650, 1962, 1996\n577 <-> 1501, 1670\n578 <-> 1504\n579 <-> 1016, 1729\n580 <-> 58\n581 <-> 346, 581\n582 <-> 749\n583 <-> 57, 613, 1219, 1941\n584 <-> 197, 1366\n585 <-> 52, 1518\n586 <-> 389, 1597\n587 <-> 227, 342\n588 <-> 1408\n589 <-> 1229\n590 <-> 143, 392, 1044, 1078\n591 <-> 448, 1289\n592 <-> 592\n593 <-> 711\n594 <-> 770\n595 <-> 290, 1169\n596 <-> 1854\n597 <-> 528, 755, 1241, 1982\n598 <-> 663, 806\n599 <-> 266\n600 <-> 1065, 1694\n601 <-> 746\n602 <-> 827\n603 <-> 90, 471, 480, 1083\n604 <-> 138, 1003\n605 <-> 1703\n606 <-> 1289\n607 <-> 1309\n608 <-> 808, 1100\n609 <-> 529, 1621, 1943\n610 <-> 194, 610\n611 <-> 1337, 1609\n612 <-> 980, 1350\n613 <-> 468, 583\n614 <-> 25, 173, 1623, 1914\n615 <-> 364, 775\n616 <-> 266, 290\n617 <-> 1514\n618 <-> 1031\n619 <-> 494, 1761\n620 <-> 1126, 1150, 1221, 1513, 1712\n621 <-> 1770\n622 <-> 144, 1761\n623 <-> 1367\n624 <-> 42, 131\n625 <-> 106, 914\n626 <-> 1849\n627 <-> 1477\n628 <-> 1379\n629 <-> 969\n630 <-> 1114\n631 <-> 218, 1935, 1979\n632 <-> 1015, 1665, 1816, 1819\n633 <-> 1326, 1628\n634 <-> 1347, 1720\n635 <-> 816\n636 <-> 636\n637 <-> 414, 1566\n638 <-> 1668\n639 <-> 544, 824, 1542\n640 <-> 87\n641 <-> 298, 1196\n642 <-> 552, 1619\n643 <-> 972\n644 <-> 814\n645 <-> 393, 1148, 1205, 1302\n646 <-> 969\n647 <-> 935, 1575\n648 <-> 648, 1424\n649 <-> 1794\n650 <-> 429, 854, 1586, 1711\n651 <-> 198, 370\n652 <-> 996, 1996\n653 <-> 150, 1919\n654 <-> 654\n655 <-> 658\n656 <-> 289\n657 <-> 978, 1954\n658 <-> 200, 655\n659 <-> 0, 825, 1258, 1792\n660 <-> 550, 1232\n661 <-> 661, 1405\n662 <-> 662\n663 <-> 598\n664 <-> 1014\n665 <-> 371, 570, 1247\n666 <-> 989, 1778\n667 <-> 227, 348, 387, 433\n668 <-> 362, 462, 703, 1091\n669 <-> 1184\n670 <-> 143\n671 <-> 671, 901\n672 <-> 171, 552\n673 <-> 1865\n674 <-> 972, 1961\n675 <-> 104, 1581, 1800\n676 <-> 334, 1995\n677 <-> 1836\n678 <-> 52\n679 <-> 169\n680 <-> 331\n681 <-> 823\n682 <-> 921\n683 <-> 1800\n684 <-> 750, 1505\n685 <-> 722, 1338, 1993\n686 <-> 269, 322\n687 <-> 989, 1178\n688 <-> 1338\n689 <-> 1535\n690 <-> 13, 690\n691 <-> 289, 721\n692 <-> 442, 1014\n693 <-> 295\n694 <-> 154, 862\n695 <-> 1552, 1735\n696 <-> 1044\n697 <-> 993, 1793, 1801\n698 <-> 698\n699 <-> 463, 1022, 1399\n700 <-> 897, 1085\n701 <-> 1782, 1987\n702 <-> 284\n703 <-> 668, 1031\n704 <-> 704, 849\n705 <-> 31\n706 <-> 1909\n707 <-> 1276\n708 <-> 1301\n709 <-> 532, 1607\n710 <-> 1094\n711 <-> 75, 593, 1251, 1297\n712 <-> 184, 712\n713 <-> 221, 388, 713\n714 <-> 198\n715 <-> 1607\n716 <-> 1110\n717 <-> 554, 1819\n718 <-> 1755\n719 <-> 472, 1942\n720 <-> 1043, 1984\n721 <-> 691, 1641\n722 <-> 685, 1360, 1679\n723 <-> 54, 174, 552, 575\n724 <-> 1466\n725 <-> 1231\n726 <-> 342, 1023\n727 <-> 24\n728 <-> 1928\n729 <-> 975, 1759\n730 <-> 28, 112, 424, 1324\n731 <-> 731\n732 <-> 76, 1698\n733 <-> 1895\n734 <-> 278\n735 <-> 571, 1440\n736 <-> 239\n737 <-> 0, 1214, 1949\n738 <-> 575, 985\n739 <-> 1467\n740 <-> 740, 1449, 1885\n741 <-> 1243\n742 <-> 845, 1588, 1676, 1956, 1980\n743 <-> 399\n744 <-> 744, 1990\n745 <-> 298, 1966\n746 <-> 601, 1294, 1601\n747 <-> 175, 1106\n748 <-> 1433, 1565, 1795\n749 <-> 366, 582\n750 <-> 548, 684, 1210\n751 <-> 751, 1967\n752 <-> 765, 1297\n753 <-> 461\n754 <-> 754\n755 <-> 597, 1089\n756 <-> 756\n757 <-> 1350, 1936\n758 <-> 444, 1044, 1279\n759 <-> 996, 1310\n760 <-> 1943\n761 <-> 217, 1333\n762 <-> 1671\n763 <-> 1761\n764 <-> 1881\n765 <-> 419, 752\n766 <-> 363, 928, 1100\n767 <-> 1220, 1986\n768 <-> 1337\n769 <-> 406, 1307\n770 <-> 467, 594, 1707\n771 <-> 1624, 1782\n772 <-> 1230\n773 <-> 514, 1266, 1305\n774 <-> 113, 1198\n775 <-> 615, 1753\n776 <-> 1307\n777 <-> 1785\n778 <-> 778, 1776\n779 <-> 992\n780 <-> 1406\n781 <-> 781, 1272\n782 <-> 93, 395, 1207\n783 <-> 1544, 1729\n784 <-> 225, 1168, 1587\n785 <-> 1050\n786 <-> 140, 162, 458, 786, 1060\n787 <-> 1210\n788 <-> 1431\n789 <-> 76\n790 <-> 164, 882\n791 <-> 166, 1688\n792 <-> 792, 946\n793 <-> 214, 561\n794 <-> 5, 1561\n795 <-> 575, 1208, 1875\n796 <-> 14, 1205, 1427\n797 <-> 336, 1006, 1352\n798 <-> 1176, 1754\n799 <-> 251, 375\n800 <-> 840, 1471\n801 <-> 801, 1246, 1897\n802 <-> 575\n803 <-> 1692\n804 <-> 565, 859, 1978\n805 <-> 288, 1447\n806 <-> 88, 598\n807 <-> 219, 1579\n808 <-> 608\n809 <-> 335, 867, 1734, 1843\n810 <-> 1279\n811 <-> 294, 1043, 1560, 1583, 1600\n812 <-> 1487, 1527\n813 <-> 42, 1405\n814 <-> 242, 644, 1175, 1638\n815 <-> 1050, 1177\n816 <-> 218, 635, 1427\n817 <-> 416\n818 <-> 1765, 1999\n819 <-> 141, 533\n820 <-> 1238, 1393\n821 <-> 504, 1755\n822 <-> 1320, 1397\n823 <-> 681, 1120, 1798\n824 <-> 639\n825 <-> 659\n826 <-> 1722, 1834\n827 <-> 193, 602, 968\n828 <-> 233, 1835\n829 <-> 465\n830 <-> 951\n831 <-> 1606\n832 <-> 1619, 1964\n833 <-> 1109\n834 <-> 7, 962, 1869\n835 <-> 141\n836 <-> 167, 875, 1145\n837 <-> 1752\n838 <-> 838, 1146\n839 <-> 1247\n840 <-> 800\n841 <-> 1337\n842 <-> 1438, 1466\n843 <-> 102\n844 <-> 284\n845 <-> 742\n846 <-> 382, 574\n847 <-> 79, 398, 409\n848 <-> 959\n849 <-> 9, 704\n850 <-> 915, 1455, 1911, 1987\n851 <-> 531, 1001\n852 <-> 434, 1172\n853 <-> 310, 1533\n854 <-> 650\n855 <-> 483, 868, 1598\n856 <-> 222, 534, 1133\n857 <-> 857, 1249, 1923\n858 <-> 412, 1540, 1575, 1607\n859 <-> 804, 1655\n860 <-> 564, 1243\n861 <-> 296, 350, 1130, 1521\n862 <-> 694\n863 <-> 272, 391, 1558\n864 <-> 309, 1071, 1227\n865 <-> 1537, 1669, 1825\n866 <-> 99, 474\n867 <-> 809, 867, 1004\n868 <-> 855, 1539\n869 <-> 1550\n870 <-> 1885\n871 <-> 49\n872 <-> 326, 378, 1552\n873 <-> 1307\n874 <-> 62, 438\n875 <-> 836, 1056, 1326\n876 <-> 876\n877 <-> 1933\n878 <-> 52, 128, 1527\n879 <-> 879\n880 <-> 1292, 1561, 1770\n881 <-> 881, 981\n882 <-> 790, 882, 932\n883 <-> 1322, 1963\n884 <-> 1842\n885 <-> 885\n886 <-> 33, 80, 1765, 1959\n887 <-> 1012\n888 <-> 276\n889 <-> 1391\n890 <-> 140\n891 <-> 114, 891\n892 <-> 160, 507\n893 <-> 236, 502, 1602\n894 <-> 1525\n895 <-> 41\n896 <-> 382\n897 <-> 700\n898 <-> 299\n899 <-> 1797\n900 <-> 1736\n901 <-> 671, 944, 1147\n902 <-> 1597, 1940\n903 <-> 1283\n904 <-> 17, 91, 557, 1169, 1764\n905 <-> 505\n906 <-> 9, 69, 942\n907 <-> 80, 211, 1348\n908 <-> 1398\n909 <-> 909\n910 <-> 1683\n911 <-> 73, 911\n912 <-> 1929\n913 <-> 12, 491, 1591\n914 <-> 169, 243, 625, 1867\n915 <-> 850, 1538\n916 <-> 83\n917 <-> 378\n918 <-> 1438\n919 <-> 1951\n920 <-> 1814\n921 <-> 360, 682, 1952\n922 <-> 922\n923 <-> 107, 1161, 1254\n924 <-> 222\n925 <-> 1691\n926 <-> 575\n927 <-> 1165, 1270\n928 <-> 368, 766\n929 <-> 1707, 1914\n930 <-> 145, 1096, 1434, 1791\n931 <-> 931, 1983\n932 <-> 882\n933 <-> 502, 1921\n934 <-> 1297\n935 <-> 647, 1319\n936 <-> 1255\n937 <-> 1210\n938 <-> 11\n939 <-> 1348, 1666\n940 <-> 1408\n941 <-> 1440\n942 <-> 906\n943 <-> 1977\n944 <-> 901\n945 <-> 126\n946 <-> 792\n947 <-> 44, 1400\n948 <-> 948, 1045\n949 <-> 967\n950 <-> 355\n951 <-> 68, 830, 969\n952 <-> 1225, 1757, 1929\n953 <-> 1534, 1726\n954 <-> 1511, 1888\n955 <-> 183, 1418\n956 <-> 1702\n957 <-> 425\n958 <-> 372\n959 <-> 848, 1512\n960 <-> 1868\n961 <-> 1509\n962 <-> 834, 1678\n963 <-> 160\n964 <-> 101, 161\n965 <-> 355\n966 <-> 1740, 1866\n967 <-> 949, 1802\n968 <-> 827, 1534\n969 <-> 629, 646, 951\n970 <-> 311\n971 <-> 35, 1670\n972 <-> 643, 674, 1449\n973 <-> 358\n974 <-> 974\n975 <-> 729, 1955\n976 <-> 1495\n977 <-> 1603, 1889\n978 <-> 337, 465, 657, 1615, 1953\n979 <-> 1801\n980 <-> 115, 612, 1239\n981 <-> 881\n982 <-> 2, 234, 1411, 1511\n983 <-> 1970\n984 <-> 227, 984\n985 <-> 191, 569, 738\n986 <-> 1926\n987 <-> 987\n988 <-> 126, 237, 1894\n989 <-> 666, 687, 1079\n990 <-> 301, 452, 1903\n991 <-> 481, 1435\n992 <-> 110, 779, 1637\n993 <-> 21, 289, 697, 1675\n994 <-> 1151, 1639\n995 <-> 1090, 1798\n996 <-> 652, 759, 1229\n997 <-> 1253, 1380, 1553\n998 <-> 350, 1812\n999 <-> 1128\n1000 <-> 1566\n1001 <-> 851\n1002 <-> 103, 1229\n1003 <-> 604, 1156, 1232, 1420\n1004 <-> 23, 867\n1005 <-> 1085, 1174, 1899\n1006 <-> 797\n1007 <-> 63, 1282\n1008 <-> 219, 1008, 1977\n1009 <-> 265, 384, 1731\n1010 <-> 275, 1438, 1474\n1011 <-> 1289\n1012 <-> 166, 235, 887, 1342, 1471\n1013 <-> 1013, 1886\n1014 <-> 664, 692, 1229\n1015 <-> 632, 1015\n1016 <-> 579\n1017 <-> 161, 528\n1018 <-> 323\n1019 <-> 1684\n1020 <-> 125\n1021 <-> 1036, 1450\n1022 <-> 699, 1600, 1962\n1023 <-> 263, 726\n1024 <-> 1295\n1025 <-> 154, 1025\n1026 <-> 268, 572, 1610\n1027 <-> 239, 1269\n1028 <-> 291, 1275\n1029 <-> 418, 479, 1957\n1030 <-> 1184\n1031 <-> 618, 703\n1032 <-> 345\n1033 <-> 61, 339, 1033\n1034 <-> 1166, 1291\n1035 <-> 31, 41\n1036 <-> 1021, 1228\n1037 <-> 28, 443, 1227\n1038 <-> 1178\n1039 <-> 94, 1578\n1040 <-> 320, 1992\n1041 <-> 1041\n1042 <-> 32, 1486\n1043 <-> 720, 811\n1044 <-> 157, 590, 696, 758, 1433, 1739\n1045 <-> 948\n1046 <-> 559, 1241\n1047 <-> 199, 1962\n1048 <-> 370\n1049 <-> 276, 1885\n1050 <-> 785, 815\n1051 <-> 1051, 1290\n1052 <-> 1165\n1053 <-> 193, 1446, 1488\n1054 <-> 1081\n1055 <-> 6\n1056 <-> 875\n1057 <-> 128, 1137\n1058 <-> 1112, 1173\n1059 <-> 1226, 1538\n1060 <-> 786\n1061 <-> 375, 1639, 1988\n1062 <-> 1748\n1063 <-> 210, 1692\n1064 <-> 444\n1065 <-> 316, 600, 1689\n1066 <-> 1709\n1067 <-> 1067\n1068 <-> 1068\n1069 <-> 350\n1070 <-> 1196\n1071 <-> 365, 864, 1827\n1072 <-> 446, 1142\n1073 <-> 1931\n1074 <-> 1529\n1075 <-> 1075\n1076 <-> 4, 1717, 1879, 1969\n1077 <-> 338\n1078 <-> 509, 590\n1079 <-> 989, 1282\n1080 <-> 139, 286\n1081 <-> 1054, 1992\n1082 <-> 1115, 1451, 1704\n1083 <-> 603, 1271\n1084 <-> 568, 1233\n1085 <-> 700, 1005, 1939\n1086 <-> 132, 1745, 1901\n1087 <-> 1519\n1088 <-> 412\n1089 <-> 755\n1090 <-> 544, 548, 995, 1768\n1091 <-> 227, 533, 668, 1141\n1092 <-> 95\n1093 <-> 1271, 1946\n1094 <-> 710, 1102\n1095 <-> 1546\n1096 <-> 930\n1097 <-> 1288\n1098 <-> 162, 1933\n1099 <-> 1456\n1100 <-> 608, 766\n1101 <-> 34, 58, 1108\n1102 <-> 1094, 1938\n1103 <-> 223, 1117\n1104 <-> 1457, 1605, 1654\n1105 <-> 1105\n1106 <-> 92, 747\n1107 <-> 1699\n1108 <-> 1101, 1201\n1109 <-> 354, 833, 1285, 1874\n1110 <-> 40, 716\n1111 <-> 190, 572, 1440\n1112 <-> 1058, 1193\n1113 <-> 1113\n1114 <-> 325, 630, 1853\n1115 <-> 260, 1082\n1116 <-> 504\n1117 <-> 380, 1103\n1118 <-> 1118\n1119 <-> 1353, 1871\n1120 <-> 823\n1121 <-> 492, 1196\n1122 <-> 1122\n1123 <-> 1725\n1124 <-> 1892\n1125 <-> 1344\n1126 <-> 620\n1127 <-> 286, 1138\n1128 <-> 999, 1268\n1129 <-> 1129\n1130 <-> 861\n1131 <-> 1874\n1132 <-> 1913\n1133 <-> 856\n1134 <-> 1185, 1767\n1135 <-> 451, 1975\n1136 <-> 65, 356, 1487\n1137 <-> 192, 1057\n1138 <-> 1127, 1782\n1139 <-> 1240\n1140 <-> 64, 390, 1385\n1141 <-> 1091\n1142 <-> 1072, 1587\n1143 <-> 155, 1143\n1144 <-> 265\n1145 <-> 836, 1401\n1146 <-> 838\n1147 <-> 901, 1483\n1148 <-> 94, 493, 645, 1167\n1149 <-> 1818\n1150 <-> 620, 1237, 1264\n1151 <-> 994, 1844\n1152 <-> 299, 1167\n1153 <-> 443, 1947\n1154 <-> 1803\n1155 <-> 489, 1163\n1156 <-> 1003, 1635, 1668\n1157 <-> 1340, 1809\n1158 <-> 231\n1159 <-> 57\n1160 <-> 1558\n1161 <-> 923\n1162 <-> 1590\n1163 <-> 1155\n1164 <-> 145, 1164, 1283\n1165 <-> 927, 1052, 1678, 1934\n1166 <-> 1034\n1167 <-> 240, 1148, 1152, 1462\n1168 <-> 784\n1169 <-> 595, 904\n1170 <-> 421\n1171 <-> 1667\n1172 <-> 852, 1195, 1323, 1444\n1173 <-> 1058, 1389\n1174 <-> 1005\n1175 <-> 814\n1176 <-> 798\n1177 <-> 258, 815, 1401\n1178 <-> 687, 1038, 1331\n1179 <-> 134, 1179\n1180 <-> 420\n1181 <-> 1181\n1182 <-> 267, 1677\n1183 <-> 100\n1184 <-> 669, 1030, 1969\n1185 <-> 498, 1134, 1673, 1750\n1186 <-> 165\n1187 <-> 1893\n1188 <-> 1236, 1365\n1189 <-> 1334, 1732\n1190 <-> 279\n1191 <-> 1620\n1192 <-> 1257\n1193 <-> 1112\n1194 <-> 1770\n1195 <-> 187, 1172, 1964\n1196 <-> 641, 1070, 1121, 1729\n1197 <-> 511, 1273, 1607\n1198 <-> 229, 774\n1199 <-> 290, 455, 1860\n1200 <-> 1901\n1201 <-> 1108\n1202 <-> 1378\n1203 <-> 1591\n1204 <-> 229\n1205 <-> 645, 796, 1250\n1206 <-> 281, 1824\n1207 <-> 782\n1208 <-> 546, 795, 1700\n1209 <-> 249\n1210 <-> 750, 787, 937\n1211 <-> 527, 1981\n1212 <-> 1212, 1369\n1213 <-> 1512\n1214 <-> 737\n1215 <-> 575\n1216 <-> 540, 1216\n1217 <-> 1312, 1340\n1218 <-> 291, 1586\n1219 <-> 583, 1554\n1220 <-> 117, 767\n1221 <-> 411, 620, 1221\n1222 <-> 1612, 1710\n1223 <-> 1223, 1374\n1224 <-> 331\n1225 <-> 315, 952, 1263\n1226 <-> 1059, 1928\n1227 <-> 864, 1037, 1823\n1228 <-> 1036, 1613\n1229 <-> 589, 996, 1002, 1014, 1906\n1230 <-> 247, 468, 772, 1821, 1837\n1231 <-> 195, 220, 725\n1232 <-> 271, 542, 660, 1003, 1232, 1606\n1233 <-> 1084, 1366, 1738\n1234 <-> 543, 1408\n1235 <-> 524\n1236 <-> 183, 1188\n1237 <-> 1150\n1238 <-> 820, 1521\n1239 <-> 980\n1240 <-> 1139, 1240\n1241 <-> 597, 1046\n1242 <-> 1646\n1243 <-> 331, 741, 860\n1244 <-> 1522, 1870\n1245 <-> 1245\n1246 <-> 801\n1247 <-> 665, 839\n1248 <-> 208, 1932\n1249 <-> 857\n1250 <-> 1205\n1251 <-> 711\n1252 <-> 1888\n1253 <-> 997, 1593\n1254 <-> 235, 520, 923, 1382\n1255 <-> 936\n1256 <-> 1718\n1257 <-> 259, 1192\n1258 <-> 659, 1917\n1259 <-> 392\n1260 <-> 1260\n1261 <-> 1261, 1376\n1262 <-> 1936\n1263 <-> 1225\n1264 <-> 1150, 1567\n1265 <-> 398\n1266 <-> 773, 1373\n1267 <-> 257, 341, 525\n1268 <-> 202, 262, 1128\n1269 <-> 1027\n1270 <-> 927\n1271 <-> 1083, 1093\n1272 <-> 781\n1273 <-> 1197\n1274 <-> 1760\n1275 <-> 1028, 1827\n1276 <-> 707, 1900\n1277 <-> 188\n1278 <-> 1463\n1279 <-> 758, 810\n1280 <-> 1920\n1281 <-> 1942\n1282 <-> 53, 1007, 1079\n1283 <-> 192, 903, 1164, 1690\n1284 <-> 512, 1584\n1285 <-> 1109, 1285\n1286 <-> 272\n1287 <-> 160\n1288 <-> 15, 1097, 1905\n1289 <-> 591, 606, 1011\n1290 <-> 1051\n1291 <-> 331, 1034\n1292 <-> 470, 880\n1293 <-> 1501, 1899\n1294 <-> 746, 1294\n1295 <-> 183, 1024\n1296 <-> 1580\n1297 <-> 711, 752, 934\n1298 <-> 1298\n1299 <-> 517, 1299\n1300 <-> 1300\n1301 <-> 708, 1301\n1302 <-> 645\n1303 <-> 1355, 1492\n1304 <-> 1361, 1746\n1305 <-> 773\n1306 <-> 334, 1325\n1307 <-> 769, 776, 873, 1333\n1308 <-> 312, 1851, 1966\n1309 <-> 607, 1459, 1496\n1310 <-> 759\n1311 <-> 1381, 1520\n1312 <-> 1217\n1313 <-> 1983\n1314 <-> 1314\n1315 <-> 1315, 1647\n1316 <-> 63\n1317 <-> 537\n1318 <-> 1973\n1319 <-> 935, 1517\n1320 <-> 531, 822\n1321 <-> 27, 1837\n1322 <-> 883, 1664\n1323 <-> 1172\n1324 <-> 730\n1325 <-> 1306\n1326 <-> 633, 875\n1327 <-> 159, 1553\n1328 <-> 1974\n1329 <-> 1939\n1330 <-> 277, 1515\n1331 <-> 1178, 1590\n1332 <-> 274, 1332\n1333 <-> 8, 761, 1307, 1333\n1334 <-> 275, 1189, 1482\n1335 <-> 1478\n1336 <-> 140\n1337 <-> 66, 611, 768, 841\n1338 <-> 176, 685, 688, 1449\n1339 <-> 1339\n1340 <-> 301, 1157, 1217, 1630\n1341 <-> 27\n1342 <-> 1012\n1343 <-> 1408\n1344 <-> 452, 1125, 1622\n1345 <-> 74, 1835\n1346 <-> 1860\n1347 <-> 634, 1428\n1348 <-> 907, 939\n1349 <-> 490, 1349\n1350 <-> 305, 612, 757\n1351 <-> 162\n1352 <-> 797, 1818\n1353 <-> 1119, 1353\n1354 <-> 1436\n1355 <-> 1303\n1356 <-> 94, 254\n1357 <-> 1419, 1664, 1680\n1358 <-> 1358\n1359 <-> 197\n1360 <-> 722, 1909\n1361 <-> 1304\n1362 <-> 82, 188, 1448\n1363 <-> 1752\n1364 <-> 501\n1365 <-> 1188, 1661\n1366 <-> 584, 1233\n1367 <-> 623, 1394, 1781\n1368 <-> 1886\n1369 <-> 1212\n1370 <-> 86, 1370\n1371 <-> 1772\n1372 <-> 191, 1473\n1373 <-> 1266\n1374 <-> 1223, 1981\n1375 <-> 366, 1375\n1376 <-> 1261, 1599\n1377 <-> 1675\n1378 <-> 30, 1202, 1406, 1845\n1379 <-> 153, 628, 1557\n1380 <-> 180, 997\n1381 <-> 1311, 1407, 1666\n1382 <-> 1254\n1383 <-> 375, 1714\n1384 <-> 392\n1385 <-> 1140, 1933\n1386 <-> 1949\n1387 <-> 439, 1387\n1388 <-> 1770\n1389 <-> 1173, 1679\n1390 <-> 123, 453, 1733\n1391 <-> 158, 249, 889, 1945\n1392 <-> 1881\n1393 <-> 820\n1394 <-> 18, 181, 355, 1367\n1395 <-> 374, 1404\n1396 <-> 515, 1860\n1397 <-> 200, 822\n1398 <-> 472, 908, 1622, 1701\n1399 <-> 699\n1400 <-> 422, 947, 1551\n1401 <-> 1145, 1177\n1402 <-> 1764\n1403 <-> 1875\n1404 <-> 1395\n1405 <-> 661, 813\n1406 <-> 780, 1378\n1407 <-> 1381\n1408 <-> 588, 940, 1234, 1343, 1603, 1865\n1409 <-> 1427\n1410 <-> 164\n1411 <-> 88, 982\n1412 <-> 1452\n1413 <-> 1707, 1766\n1414 <-> 121\n1415 <-> 318, 1415, 1612\n1416 <-> 1416\n1417 <-> 232\n1418 <-> 361, 955, 1418, 1737\n1419 <-> 473, 1357\n1420 <-> 1003\n1421 <-> 122, 1973\n1422 <-> 512, 1870\n1423 <-> 248\n1424 <-> 648\n1425 <-> 1425\n1426 <-> 224, 1946\n1427 <-> 796, 816, 1409\n1428 <-> 1347\n1429 <-> 1810, 1862\n1430 <-> 1430\n1431 <-> 788, 1488\n1432 <-> 259, 1432\n1433 <-> 1, 501, 748, 1044\n1434 <-> 930\n1435 <-> 991\n1436 <-> 1354, 1436\n1437 <-> 253\n1438 <-> 522, 558, 842, 918, 1010\n1439 <-> 1649\n1440 <-> 735, 941, 1111\n1441 <-> 1707\n1442 <-> 305, 1930\n1443 <-> 357\n1444 <-> 1172\n1445 <-> 162, 270, 1636\n1446 <-> 1053\n1447 <-> 220, 292, 805\n1448 <-> 1362\n1449 <-> 740, 972, 1338\n1450 <-> 100, 237, 1021\n1451 <-> 5, 1082\n1452 <-> 1412, 1572\n1453 <-> 556\n1454 <-> 1454\n1455 <-> 850\n1456 <-> 125, 268, 1099\n1457 <-> 240, 1104\n1458 <-> 18, 1576\n1459 <-> 1309, 1503\n1460 <-> 84, 371\n1461 <-> 1747, 1879\n1462 <-> 3, 1167, 1807\n1463 <-> 377, 1278, 1499, 1971\n1464 <-> 1908\n1465 <-> 72\n1466 <-> 724, 842\n1467 <-> 410, 739\n1468 <-> 67\n1469 <-> 155, 1652\n1470 <-> 268, 469, 486\n1471 <-> 800, 1012, 1471, 1510\n1472 <-> 39\n1473 <-> 1372\n1474 <-> 1010, 1915\n1475 <-> 467, 1658\n1476 <-> 1476\n1477 <-> 627, 1477\n1478 <-> 427, 1335, 1907\n1479 <-> 348\n1480 <-> 1480\n1481 <-> 1802\n1482 <-> 1334\n1483 <-> 1147, 1524\n1484 <-> 42, 525\n1485 <-> 539\n1486 <-> 148, 1042, 1549\n1487 <-> 812, 1136\n1488 <-> 1053, 1431\n1489 <-> 1708, 1931\n1490 <-> 1807\n1491 <-> 431\n1492 <-> 156, 1303\n1493 <-> 1493\n1494 <-> 1853\n1495 <-> 446, 976\n1496 <-> 395, 1309\n1497 <-> 195, 396, 1974\n1498 <-> 490, 1626\n1499 <-> 76, 1463\n1500 <-> 1722\n1501 <-> 577, 1293\n1502 <-> 115\n1503 <-> 1459, 1507\n1504 <-> 210, 261, 578, 1521, 1718\n1505 <-> 230, 426, 684, 1634\n1506 <-> 382, 1921\n1507 <-> 1503\n1508 <-> 162, 1667\n1509 <-> 961, 1509\n1510 <-> 1471\n1511 <-> 954, 982, 1582\n1512 <-> 476, 959, 1213, 1703\n1513 <-> 353, 620\n1514 <-> 617, 1546\n1515 <-> 1330, 1595\n1516 <-> 1946\n1517 <-> 42, 1319\n1518 <-> 585\n1519 <-> 32, 1087\n1520 <-> 1311\n1521 <-> 861, 1238, 1504\n1522 <-> 1244, 1653\n1523 <-> 1523\n1524 <-> 1483\n1525 <-> 894, 1525\n1526 <-> 1924\n1527 <-> 812, 878\n1528 <-> 494\n1529 <-> 59, 1074\n1530 <-> 1530\n1531 <-> 34, 1531\n1532 <-> 1638\n1533 <-> 853\n1534 <-> 953, 968\n1535 <-> 108, 689\n1536 <-> 1589\n1537 <-> 394, 865\n1538 <-> 915, 1059\n1539 <-> 868\n1540 <-> 858, 1745\n1541 <-> 221\n1542 <-> 639\n1543 <-> 1746\n1544 <-> 783, 1544\n1545 <-> 236\n1546 <-> 66, 1095, 1514\n1547 <-> 1628\n1548 <-> 1548\n1549 <-> 433, 1486\n1550 <-> 869, 1550\n1551 <-> 1400, 1787\n1552 <-> 695, 872\n1553 <-> 997, 1327\n1554 <-> 1219\n1555 <-> 97\n1556 <-> 1840\n1557 <-> 7, 1379\n1558 <-> 126, 863, 1160\n1559 <-> 336, 373\n1560 <-> 343, 811\n1561 <-> 278, 794, 880, 1561\n1562 <-> 127, 518, 1562\n1563 <-> 406\n1564 <-> 81, 399\n1565 <-> 748\n1566 <-> 637, 1000\n1567 <-> 1264\n1568 <-> 1568\n1569 <-> 214\n1570 <-> 561, 1849\n1571 <-> 282\n1572 <-> 44, 130, 178, 1452\n1573 <-> 107\n1574 <-> 299\n1575 <-> 647, 858\n1576 <-> 1458, 1633\n1577 <-> 1833, 1939\n1578 <-> 1039, 1802\n1579 <-> 807, 1853\n1580 <-> 1296, 1580, 1907\n1581 <-> 675\n1582 <-> 1511, 1605, 1756\n1583 <-> 811\n1584 <-> 1284\n1585 <-> 1817\n1586 <-> 168, 650, 1218\n1587 <-> 784, 1142\n1588 <-> 742\n1589 <-> 45, 209, 413, 1536\n1590 <-> 25, 201, 1162, 1331\n1591 <-> 913, 1203\n1592 <-> 1820\n1593 <-> 1253\n1594 <-> 231\n1595 <-> 149, 1515\n1596 <-> 57, 1976\n1597 <-> 586, 902\n1598 <-> 855\n1599 <-> 1376\n1600 <-> 811, 1022\n1601 <-> 284, 746\n1602 <-> 386, 893\n1603 <-> 135, 977, 1408\n1604 <-> 452, 513\n1605 <-> 1104, 1582\n1606 <-> 831, 1232\n1607 <-> 709, 715, 858, 1197\n1608 <-> 1793\n1609 <-> 611, 1808\n1610 <-> 1026, 1964\n1611 <-> 304\n1612 <-> 1222, 1415, 1769\n1613 <-> 1228, 1847\n1614 <-> 273\n1615 <-> 978\n1616 <-> 156, 355\n1617 <-> 177\n1618 <-> 1618\n1619 <-> 323, 642, 832\n1620 <-> 113, 349, 1191, 1746\n1621 <-> 609\n1622 <-> 519, 1344, 1398\n1623 <-> 614\n1624 <-> 771, 1989\n1625 <-> 1625\n1626 <-> 1498\n1627 <-> 186, 315\n1628 <-> 633, 1547\n1629 <-> 1706\n1630 <-> 1340\n1631 <-> 297, 1800\n1632 <-> 1806\n1633 <-> 1576\n1634 <-> 1505\n1635 <-> 1156\n1636 <-> 1445\n1637 <-> 76, 992, 1920, 1995\n1638 <-> 814, 1532\n1639 <-> 994, 1061\n1640 <-> 1640\n1641 <-> 721\n1642 <-> 1739, 1945\n1643 <-> 1643\n1644 <-> 1644, 1908\n1645 <-> 125, 303\n1646 <-> 1242, 1646\n1647 <-> 1315\n1648 <-> 1957\n1649 <-> 97, 1439, 1783\n1650 <-> 576\n1651 <-> 1651\n1652 <-> 35, 1469\n1653 <-> 1522\n1654 <-> 1104\n1655 <-> 460, 859\n1656 <-> 362\n1657 <-> 142, 181\n1658 <-> 428, 1475\n1659 <-> 470\n1660 <-> 247\n1661 <-> 1365\n1662 <-> 165, 359\n1663 <-> 198\n1664 <-> 1322, 1357\n1665 <-> 447, 632\n1666 <-> 939, 1381\n1667 <-> 1171, 1508\n1668 <-> 638, 1156\n1669 <-> 476, 865, 1699\n1670 <-> 577, 971\n1671 <-> 119, 417, 762, 1804\n1672 <-> 50, 1672, 1687\n1673 <-> 63, 333, 1185\n1674 <-> 1674\n1675 <-> 993, 1377\n1676 <-> 742, 1813\n1677 <-> 233, 1182, 1846\n1678 <-> 962, 1165\n1679 <-> 722, 1389\n1680 <-> 82, 1357\n1681 <-> 1681, 1702, 1895\n1682 <-> 372\n1683 <-> 203, 341, 910\n1684 <-> 1019, 1948\n1685 <-> 1982\n1686 <-> 371\n1687 <-> 1672\n1688 <-> 791\n1689 <-> 1065, 1689\n1690 <-> 1283\n1691 <-> 925, 1878\n1692 <-> 803, 1063, 1732\n1693 <-> 487\n1694 <-> 600\n1695 <-> 1695, 1794\n1696 <-> 1696\n1697 <-> 1802\n1698 <-> 357, 732\n1699 <-> 1107, 1669\n1700 <-> 1208\n1701 <-> 1398\n1702 <-> 956, 1681\n1703 <-> 605, 1512\n1704 <-> 1082\n1705 <-> 343\n1706 <-> 1629, 1775\n1707 <-> 770, 929, 1413, 1441\n1708 <-> 367, 1489\n1709 <-> 1066, 1862\n1710 <-> 1222\n1711 <-> 650\n1712 <-> 620\n1713 <-> 1713\n1714 <-> 1383\n1715 <-> 1911\n1716 <-> 510, 1716\n1717 <-> 1076\n1718 <-> 1256, 1504\n1719 <-> 238\n1720 <-> 549, 634, 1720\n1721 <-> 15\n1722 <-> 826, 1500\n1723 <-> 1880\n1724 <-> 177, 1724\n1725 <-> 514, 1123, 1797\n1726 <-> 482, 953\n1727 <-> 1914\n1728 <-> 216, 1982\n1729 <-> 579, 783, 1196, 1800\n1730 <-> 1730, 1972\n1731 <-> 1009\n1732 <-> 89, 1189, 1692\n1733 <-> 1390\n1734 <-> 809\n1735 <-> 455, 695\n1736 <-> 900, 1736\n1737 <-> 530, 1418\n1738 <-> 1233, 1888\n1739 <-> 1044, 1642, 1834\n1740 <-> 966\n1741 <-> 145, 253\n1742 <-> 112\n1743 <-> 1980\n1744 <-> 139\n1745 <-> 1086, 1540\n1746 <-> 1304, 1543, 1620\n1747 <-> 117, 354, 1461\n1748 <-> 45, 273, 332, 1062\n1749 <-> 249\n1750 <-> 124, 1185\n1751 <-> 1835\n1752 <-> 104, 837, 1363\n1753 <-> 245, 775\n1754 <-> 798, 1876\n1755 <-> 453, 718, 821\n1756 <-> 1582\n1757 <-> 952, 1772, 1841\n1758 <-> 1758, 1830, 1878\n1759 <-> 729\n1760 <-> 104, 1274\n1761 <-> 224, 619, 622, 763, 1868\n1762 <-> 1762\n1763 <-> 537\n1764 <-> 904, 1402\n1765 <-> 521, 818, 886, 1782\n1766 <-> 421, 1413\n1767 <-> 1134\n1768 <-> 202, 1090\n1769 <-> 1612\n1770 <-> 621, 880, 1194, 1388\n1771 <-> 293\n1772 <-> 1371, 1757\n1773 <-> 19, 375\n1774 <-> 1774\n1775 <-> 1706, 1915\n1776 <-> 778\n1777 <-> 460, 1918\n1778 <-> 666\n1779 <-> 523, 1846\n1780 <-> 427\n1781 <-> 1367\n1782 <-> 701, 771, 1138, 1765\n1783 <-> 482, 1649, 1783\n1784 <-> 1784, 1872\n1785 <-> 777, 1785\n1786 <-> 1858\n1787 <-> 1551\n1788 <-> 59, 384, 1865\n1789 <-> 111\n1790 <-> 294, 407, 1969\n1791 <-> 238, 930\n1792 <-> 659\n1793 <-> 489, 538, 697, 1608\n1794 <-> 206, 649, 1695\n1795 <-> 314, 748\n1796 <-> 1796\n1797 <-> 899, 1725, 1797\n1798 <-> 823, 995\n1799 <-> 308, 1799\n1800 <-> 675, 683, 1631, 1729\n1801 <-> 697, 979, 1858\n1802 <-> 967, 1481, 1578, 1697\n1803 <-> 1154, 1803\n1804 <-> 1671\n1805 <-> 199\n1806 <-> 264, 1632, 1806\n1807 <-> 1462, 1490\n1808 <-> 1609, 1808\n1809 <-> 1157\n1810 <-> 1429\n1811 <-> 125\n1812 <-> 998\n1813 <-> 1676\n1814 <-> 358, 485, 920, 1814\n1815 <-> 211\n1816 <-> 632\n1817 <-> 513, 1585\n1818 <-> 1149, 1352\n1819 <-> 632, 717\n1820 <-> 465, 1592\n1821 <-> 46, 257, 1230\n1822 <-> 492\n1823 <-> 1227, 1891\n1824 <-> 247, 1206\n1825 <-> 182, 865\n1826 <-> 501\n1827 <-> 3, 1071, 1275\n1828 <-> 29\n1829 <-> 139\n1830 <-> 1758\n1831 <-> 542\n1832 <-> 1862\n1833 <-> 1577\n1834 <-> 151, 826, 1739, 1882\n1835 <-> 828, 1345, 1751, 1835\n1836 <-> 45, 677\n1837 <-> 1230, 1321\n1838 <-> 165, 1856\n1839 <-> 1842, 1953\n1840 <-> 183, 1556\n1841 <-> 554, 1757\n1842 <-> 884, 1839\n1843 <-> 464, 809\n1844 <-> 291, 570, 1151\n1845 <-> 1378\n1846 <-> 1677, 1779\n1847 <-> 374, 1613\n1848 <-> 204, 412\n1849 <-> 626, 1570\n1850 <-> 1850\n1851 <-> 1308\n1852 <-> 228, 455\n1853 <-> 454, 1114, 1494, 1579\n1854 <-> 268, 596\n1855 <-> 1855\n1856 <-> 1838\n1857 <-> 61\n1858 <-> 336, 1786, 1801, 1937\n1859 <-> 1908\n1860 <-> 1199, 1346, 1396\n1861 <-> 1861\n1862 <-> 1429, 1709, 1832, 1900\n1863 <-> 9\n1864 <-> 41, 1975\n1865 <-> 673, 1408, 1788\n1866 <-> 966, 1873\n1867 <-> 914\n1868 <-> 330, 520, 960, 1761\n1869 <-> 2, 834\n1870 <-> 1244, 1422\n1871 <-> 1119\n1872 <-> 1784\n1873 <-> 242, 1866\n1874 <-> 337, 1109, 1131\n1875 <-> 795, 1403\n1876 <-> 258, 1754\n1877 <-> 34\n1878 <-> 481, 1691, 1758\n1879 <-> 1076, 1461\n1880 <-> 147, 410, 1723\n1881 <-> 520, 764, 1392, 1955\n1882 <-> 1834\n1883 <-> 319, 514\n1884 <-> 1969\n1885 <-> 740, 870, 1049\n1886 <-> 1013, 1368\n1887 <-> 1887\n1888 <-> 954, 1252, 1738\n1889 <-> 977\n1890 <-> 446, 530\n1891 <-> 1823\n1892 <-> 1124, 1892\n1893 <-> 395, 1187, 1893\n1894 <-> 988\n1895 <-> 733, 1681\n1896 <-> 11\n1897 <-> 77, 801\n1898 <-> 545, 1898\n1899 <-> 1005, 1293\n1900 <-> 511, 1276, 1862\n1901 <-> 1086, 1200\n1902 <-> 410\n1903 <-> 315, 990\n1904 <-> 422\n1905 <-> 1288, 1905\n1906 <-> 1229\n1907 <-> 1478, 1580\n1908 <-> 1464, 1644, 1859\n1909 <-> 706, 1360\n1910 <-> 569\n1911 <-> 310, 850, 1715\n1912 <-> 1912\n1913 <-> 279, 1132\n1914 <-> 614, 929, 1727\n1915 <-> 1474, 1775\n1916 <-> 1916\n1917 <-> 1258\n1918 <-> 382, 1777\n1919 <-> 41, 653\n1920 <-> 1280, 1637\n1921 <-> 933, 1506\n1922 <-> 167\n1923 <-> 857\n1924 <-> 222, 1526, 1924\n1925 <-> 231, 367, 1925\n1926 <-> 385, 422, 986\n1927 <-> 321, 513, 553\n1928 <-> 728, 1226\n1929 <-> 912, 952, 1965\n1930 <-> 1442\n1931 <-> 1073, 1489\n1932 <-> 1248, 1932\n1933 <-> 499, 877, 1098, 1385\n1934 <-> 1165\n1935 <-> 37, 631\n1936 <-> 283, 757, 1262\n1937 <-> 1858\n1938 <-> 109, 555, 1102\n1939 <-> 1085, 1329, 1577\n1940 <-> 902, 1940\n1941 <-> 364, 583\n1942 <-> 719, 1281\n1943 <-> 609, 760\n1944 <-> 230\n1945 <-> 1391, 1642\n1946 <-> 1093, 1426, 1516\n1947 <-> 1153\n1948 <-> 212, 296, 500, 1684\n1949 <-> 152, 737, 1386\n1950 <-> 133, 296\n1951 <-> 919, 1951\n1952 <-> 921\n1953 <-> 978, 1839\n1954 <-> 657\n1955 <-> 975, 1881\n1956 <-> 55, 71, 235, 742\n1957 <-> 1029, 1648\n1958 <-> 1958\n1959 <-> 886\n1960 <-> 196, 301, 495\n1961 <-> 674\n1962 <-> 576, 1022, 1047\n1963 <-> 883\n1964 <-> 832, 1195, 1610\n1965 <-> 1929\n1966 <-> 745, 1308\n1967 <-> 751\n1968 <-> 524, 1968\n1969 <-> 1076, 1184, 1790, 1884\n1970 <-> 983, 1970\n1971 <-> 1463\n1972 <-> 1730\n1973 <-> 457, 1318, 1421\n1974 <-> 1328, 1497, 1974\n1975 <-> 170, 1135, 1864\n1976 <-> 1596\n1977 <-> 70, 943, 1008\n1978 <-> 804\n1979 <-> 631\n1980 <-> 522, 742, 1743\n1981 <-> 1211, 1374\n1982 <-> 597, 1685, 1728\n1983 <-> 30, 931, 1313\n1984 <-> 109, 720\n1985 <-> 111\n1986 <-> 767\n1987 <-> 418, 701, 850\n1988 <-> 1061\n1989 <-> 1624\n1990 <-> 744\n1991 <-> 316\n1992 <-> 346, 562, 1040, 1081\n1993 <-> 685\n1994 <-> 573\n1995 <-> 381, 676, 1637, 1995\n1996 <-> 576, 652\n1997 <-> 554\n1998 <-> 168\n1999 <-> 818"

var examples = []Example{
	{input: "0 <-> 2\n1 <-> 1\n2 <-> 0, 3, 4\n3 <-> 2, 4\n4 <-> 2, 3, 6\n5 <-> 6\n6 <-> 4, 5", part1: 6, part2: 2},
	{input: personalData, part1: 115, part2: 221},
}

func TestKnots(t *testing.T) {
	for i := 0; i < len(examples); i++ {
		example := examples[i]
		resPart1, resPart2 := Groups(example.input)
		if resPart1 != example.part1 || resPart2 != example.part2 {
			t.Error(fmt.Sprintf("Expected %d, %d but got %d, %d", example.part1, example.part2, resPart1, resPart2))
		}
	}
}
