void expectf(float a, float b);
void expectd(double a, double b);

int main() {
    printf("Testing numeric constants ... ");

    expect(1, 0x1);
    expect(17, 0x11);
    expect(511, 0777);

    expect(3, 3L);
    expect(3, 3LL);
    expect(3, 3UL);
    expect(3, 3LU);
    expect(3, 3ULL);
    expect(3, 3LU);
    expect(3, 3LLU);

    expectd(55.3, 55.3);
    expectd(200, 2e2);
    expectd(0x0.DE488631p8, 0xDE.488631);

    expect(4, sizeof(5));
    expect(8, sizeof(5L));
    expect(4, sizeof(3.0f));
    expect(8, sizeof(3.0));

    printf("OK\n");
    return 0;
}