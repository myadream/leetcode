package com.myadream.code.leetcode.easy._0933;

import com.myadream.code.leetcode.common.DataSet;
import com.myadream.code.leetcode.common.DataSetControl;
import kotlin.jvm.functions.Function1;
import org.jetbrains.annotations.NotNull;
import org.testng.annotations.Test;

import java.util.ArrayList;

import static org.junit.jupiter.api.Assertions.assertEquals;

@Test(groups = {"tags.queue", "tags.dataStream", "tags.design", "difficulty-easy"})
class RecentCounterTest extends DataSetControl {

    @NotNull
    @Override
    public ArrayList<DataSet> buildDataSet() {
        ArrayList<DataSet> dataSets = new ArrayList<>();

        dataSets.add(
                new DataSet(
                        new int[][]{
                                new int[]{1},
                                new int[]{100},
                                new int[]{3001},
                                new int[]{3002},
                        },
                        new int[]{
                                1, 2, 3, 3
                        }
                )
        );

        return dataSets;
    }

    @NotNull
    @Override
    public ArrayList<Function1<DataSet, Object>> buildImpl() {
        ArrayList<Function1<DataSet, Object>> impls = new ArrayList<>();

        impls.add(
                dataSets -> {
                    RecentCounter recentCounter = new RecentCounter();
                    int[][] sample = (int[][]) dataSets.getSample();
                    int[] target = (int[]) dataSets.getTarget();

                    assertEquals(target[0], recentCounter.ping(sample[0][0]));
                    assertEquals(target[1], recentCounter.ping(sample[1][0]));
                    assertEquals(target[2], recentCounter.ping(sample[2][0]));
                    assertEquals(target[3], recentCounter.ping(sample[3][0]));

                    return true;
                }
        );

        impls.add(
                dataSets -> {
                    RecentCounter2 recentCounter = new RecentCounter2();
                    int[][] sample = (int[][]) dataSets.getSample();
                    int[] target = (int[]) dataSets.getTarget();

                    assertEquals(target[0], recentCounter.ping(sample[0][0]));
                    assertEquals(target[1], recentCounter.ping(sample[1][0]));
                    assertEquals(target[2], recentCounter.ping(sample[2][0]));
                    assertEquals(target[3], recentCounter.ping(sample[3][0]));

                    return true;
                }
        );

        return impls;
    }

}