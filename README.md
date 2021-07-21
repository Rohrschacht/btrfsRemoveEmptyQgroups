# btrfsRemoveEmptyQgroups

## About

btrfsRemoveEmptyQgroups deletes all qgroups that are empty, that are left over
after deleting a subvolume. Quite some of them can accumulate over time if one
uses an automation for creating and deleting snapshots for example.

## Usage

Output of `btrfs qgroup show /mnt/btrfs` before:

```
qgroupid         rfer         excl
--------         ----         ----
0/5          16.00KiB     16.00KiB
0/1098       41.80GiB      4.81GiB
0/1100       16.00KiB     16.00KiB
0/1501          0.00B        0.00B
0/1502          0.00B        0.00B
0/1503          0.00B        0.00B
0/1504          0.00B        0.00B
0/1505          0.00B        0.00B
0/1521       85.95GiB      5.20MiB
0/1575       42.16GiB      2.64GiB
0/1592        1.25GiB      1.25GiB
0/1642       43.08GiB      1.66GiB
0/1724       39.70GiB    975.56MiB
0/1745       40.61GiB    746.57MiB
0/1750       16.00KiB     16.00KiB
0/1759       39.92GiB     94.48MiB
0/1760       81.50GiB      2.33GiB
0/1765          0.00B        0.00B
0/1770       39.59GiB     84.02MiB
0/1824       40.12GiB    258.02MiB
0/1873       40.57GiB    361.68MiB
0/1874       84.61GiB    236.67MiB
0/1879       85.95GiB     40.02MiB
0/1880       85.95GiB      4.89MiB
```

Run the tool:

```sh
$ sudo btrfsRemoveEmptyQgroups /mnt/btrfs
```

Output of `btrfs qgroup show /mnt/btrfs` after:

```
qgroupid         rfer         excl
--------         ----         ----
0/5          16.00KiB     16.00KiB
0/1098       41.80GiB      4.81GiB
0/1100       16.00KiB     16.00KiB
0/1521       85.95GiB      5.20MiB
0/1575       42.16GiB      2.64GiB
0/1592        1.25GiB      1.25GiB
0/1642       43.08GiB      1.66GiB
0/1724       39.70GiB    975.56MiB
0/1745       40.61GiB    746.57MiB
0/1750       16.00KiB     16.00KiB
0/1759       39.92GiB     94.48MiB
0/1760       81.50GiB      2.33GiB
0/1770       39.59GiB     84.02MiB
0/1824       40.12GiB    258.02MiB
0/1873       40.57GiB    361.68MiB
0/1874       84.61GiB    236.67MiB
0/1879       85.95GiB     40.02MiB
0/1880       85.95GiB      4.89MiB
```
