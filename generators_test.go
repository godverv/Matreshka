package matreshka

//func Test_GenerateGoConfigKeys(t *testing.T) {
//	t.Parallel()
//
//	c, err := ParseConfig(fullConfig)
//	require.NoError(t, err)
//
//	expected := []env.EnvVal{
//		{
//			Name:  "matreshka_bool",
//			Value: true,
//		},
//		{
//			Name:  "matreshka_duration",
//			Value: "10s",
//		},
//		{
//			Name:  "matreshka_int",
//			Value: 1,
//		},
//		{
//			Name:  "matreshka_string",
//			Value: "not so basic 🤡 string",
//		},
//	}
//	expected = append(expected, getGRPCServerEnvs()...)
//	expected = append(expected, getRestServerEnvs()...)
//	expected = append(expected, getPostgresClientEnvs()...)
//	expected = append(expected, getRedisClientEnvs()...)
//	expected = append(expected, getTelegramClientEnvs()...)
//	expected = append(expected, getGRPCClientEnvs()...)
//
//	res, err := GenerateKeys(c)
//	require.NoError(t, err)
//
//	sort.Slice(res, func(i, j int) bool {
//		return res[i].Name < res[j].Name
//	})
//	sort.Slice(expected, func(i, j int) bool {
//		return expected[i].Name < expected[j].Name
//	})
//
//	require.Equal(t, expected, res)
//}
